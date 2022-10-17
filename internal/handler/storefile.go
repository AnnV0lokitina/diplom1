package handler

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
	"time"
)

const maxFileSize = 1 << 20

// reqInfo Keep information about file and user.
type reqInfo struct {
	Time      time.Time
	SessionID string
}

// saveFile Save file to store.
func (h *Handler) saveFile(ctx context.Context, sessionID string, time time.Time, r io.Reader) error {
	err := h.service.StoreFile(ctx, sessionID, time, r)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot save file to the store: %v", err))
	}
	return nil
}

// readStream Read file from stream.
func (h *Handler) readStream(stream pb.SecureStorage_StoreFileServer, w io.Writer, ch chan *reqInfo) (int, error) {
	var info *reqInfo
	firstChunk := true
	fileSize := 0

	for {
		log.Print("waiting to receive more data")
		req, err := stream.Recv()

		if err == io.EOF {
			log.Print("no more data")
			break
		}

		if err != nil {
			return fileSize, logError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}

		if firstChunk {
			firstChunk = false
			info = &reqInfo{
				SessionID: req.GetSession(),
				Time:      req.GetInfo().GetTime().AsTime(),
			}
			ch <- info
		} else {
			if info.SessionID != req.GetSession() {
				return fileSize, logError(status.Error(codes.Internal, "error in stream"))
			}
		}

		chunk := req.GetContent()
		size, err := w.Write(chunk)
		if err != nil {
			return fileSize, logError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
		}

		fileSize += size
		if fileSize > maxFileSize {
			return fileSize, logError(status.Error(codes.Internal, "to big file"))
		}
	}
	return fileSize, nil
}

// StoreFile Processed stream witch send new file to storage.
func (h *Handler) StoreFile(stream pb.SecureStorage_StoreFileServer) error {
	infoCh := make(chan *reqInfo)
	sizeCh := make(chan uint32, 1)
	ctx := context.Background()
	g, _ := errgroup.WithContext(ctx)
	r, w := io.Pipe()

	g.Go(func() error {
		log.Info("start read info from stream")
		defer w.Close()
		size, err := h.readStream(stream, w, infoCh)
		if err != nil {
			return err
		}
		sizeCh <- uint32(size)
		return nil
	})
	g.Go(func() error {
		log.Info("start write info to file")
		defer r.Close()
		info := <-infoCh
		return h.saveFile(ctx, info.SessionID, info.Time, r)
	})

	if err := g.Wait(); err != nil {
		log.Errorf("store file is failed - error: %s", err)
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
	res := &pb.StoreFileResponse{
		Size: <-sizeCh,
	}
	log.Infof("saved file size: %d", res.Size)
	err := stream.SendAndClose(res)
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot send response: %v", err))
	}

	log.Info("stream closed")
	return nil
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
