package grpctest

import (
	"context"
	"errors"
	labelError "github.com/AnnV0lokitina/diplom1/pkg/error"
	pb "github.com/AnnV0lokitina/diplom1/proto"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

// readFile Read file from store.
func (h *Handler) readFile(ctx context.Context, sessionID string, fileInfo *pb.FileInfo, w io.Writer) error {
	log.Infof("restore file updated at %s", fileInfo.GetTime().String())
	if sessionID != CorrectSession {
		if sessionID == WithErrorSession {
			return errors.New("error")
		}
		return labelError.NewLabelError(labelError.TypeUnauthorized, errors.New("user unauthorized"))
	}
	if TestFileDate.Before(fileInfo.GetTime().AsTime()) {
		log.Info("stored file is old")
		return labelError.NewLabelError(labelError.TypeUpgradeRequired, errors.New("stored file is old"))
	}
	err := h.ReadTestFile(w)
	if err != nil {
		return err
	}
	return nil
}

// sendFile Send file to user.
func (h *Handler) sendFile(stream pb.SecureStorage_RestoreFileServer, r io.Reader) error {
	buf := make([]byte, 16)
	for {
		log.Println("start send file")
		n, err := r.Read(buf)
		if err == io.EOF || n == 0 {
			log.Println("file is sent")
			break
		}
		resp := &pb.RestoreFileResponse{
			Content: buf[:n],
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
	return nil
}

// RestoreFile Restores a file from the user's storage.
func (h *Handler) RestoreFile(in *pb.RestoreFileRequest, stream pb.SecureStorage_RestoreFileServer) error {
	log.Info("start restoring file")
	sessionID := in.GetSession()
	fileInfo := in.GetInfo()
	ctx := context.Background()

	r, w := io.Pipe()
	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() error {
		log.Info("start read info from file")
		defer w.Close()
		return h.readFile(ctx, sessionID, fileInfo, w)
	})
	g.Go(func() error {
		log.Info("start send file to user")
		defer r.Close()
		return h.sendFile(stream, r)
	})
	if err := g.Wait(); err != nil {
		log.Errorf("restore file is failed - error: %s", err)
		var labelErr *labelError.LabelError
		if errors.As(err, &labelErr) {
			if labelErr.Label == labelError.TypeUnauthorized {
				log.Info("upgrade required")
				return status.Error(codes.PermissionDenied, "User unauthorized")
			}
			if labelErr.Label == labelError.TypeUpgradeRequired {
				log.Info("upgrade required")
				return status.Error(codes.FailedPrecondition, "Upgrade required")
			}
		}
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}
