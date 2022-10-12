package handler

import (
	pb "github.com/AnnV0lokitina/diplom1/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

const maxFileSize = 1 << 20

func (h *Handler) StoreFile(stream pb.SecureStorage_StoreFileServer) error {
	var (
		fileType string
		fileName string
		session  string
		time     *timestamp.Timestamp
	)

	httpShutdownCh := make(chan struct{})

	fileSize := 0
	r, w := io.Pipe()
	firstChunk := true

	for {
		log.Print("waiting to receive more data")

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			w.Close()
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}

		tmpFileType := req.GetInfo().GetType()
		tmpFileName := req.GetInfo().GetName()
		tmpTime := req.GetInfo().GetTime()
		tmpSession := req.GetSession()
		chunk := req.GetContent()

		if firstChunk {
			firstChunk = false
			fileType = tmpFileType
			fileName = tmpFileName
			time = tmpTime
			session = tmpSession

			go func() {
				err := h.service.StoreFile(session, fileType, fileName, r, time.String())
				if err != nil {
					logError(status.Errorf(codes.Internal, "cannot save file to the store: %v", err))
				}
				httpShutdownCh <- struct{}{}
			}()
		} else {
			if fileType != tmpFileType ||
				fileName != tmpFileName ||
				session != tmpSession {
				return logError(status.Error(codes.Internal, "error in stream"))
			}
		}

		size, err := w.Write(chunk)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
		}

		fileSize += size
		if fileSize > maxFileSize {
			return logError(status.Error(codes.Internal, "to big file"))
		}
	}
	<-httpShutdownCh
	res := &pb.StoreFileResponse{
		Size: uint32(fileSize),
	}

	err := stream.SendAndClose(res)
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot send response: %v", err))
	}

	log.Printf("saved file size: %d", fileSize)
	return nil
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
