package external

import (
	"context"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	pb "github.com/AnnV0lokitina/diplom1/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
)

type ExtConnection struct {
	target   string
	filePath string
}

func NewExtConnection(target string, filePath string) *ExtConnection {
	return &ExtConnection{
		target:   target,
		filePath: filePath,
	}
}

// createConnection create connection with server.
func (ec *ExtConnection) createConnection() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(ec.target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Errorf("error create grpc connection: %s", err)
		return nil, NewError(err, ErrorTypeConnection)
	}
	return conn, nil
}

func (ec *ExtConnection) Register(ctx context.Context, login string, password string) (string, error) {
	conn, err := ec.createConnection()
	if err != nil {
		log.Errorf("error create grpc connection: %s", err)
		return "", err
	}
	defer conn.Close()
	c := pb.NewSecureStorageClient(conn)
	req := &pb.UserRegRequest{
		Login:    login,
		Password: password,
	}
	resp, err := c.Register(ctx, req)
	if err != nil {
		log.Errorf("error register user: %s", err)
		return "", NewError(err, ErrorTypeRegister)
	}
	return resp.GetSession(), nil
}

func (ec *ExtConnection) Login(ctx context.Context, login string, password string) (string, error) {
	conn, err := ec.createConnection()
	if err != nil {
		log.Errorf("error create grpc connection: %s", err)
		return "", err
	}
	defer conn.Close()
	c := pb.NewSecureStorageClient(conn)
	req := &pb.UserRegRequest{
		Login:    login,
		Password: password,
	}
	resp, err := c.Login(ctx, req)
	if err != nil {
		log.Errorf("error login user: %s", err)
		return "", NewError(err, ErrorTypeLogin)
	}
	return resp.GetSession(), nil
}

func (ec *ExtConnection) StoreInfo(ctx context.Context, session string, reader io.Reader, info *entity.FileInfo) error {
	conn, err := ec.createConnection()
	if err != nil {
		log.Errorf("error create grpc connection: %s", err)
		return err
	}
	defer conn.Close()
	c := pb.NewSecureStorageClient(conn)
	stream, err := c.StoreFile(ctx)
	if err != nil {
		log.Errorf("error store file: %s", err)
		return NewError(err, ErrorTypeSaveToStorage)
	}
	buf := make([]byte, 16)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		part := pb.StoreFileRequest{
			Info: &pb.FileInfo{
				Time: timestamppb.New(info.UpdateTime),
			},
			Content: buf[:n],
			Session: session,
		}
		if err := stream.Send(&part); err != nil {
			log.Errorf("error send to stream: %s", err)
			return NewError(err, ErrorTypeSaveToStorage)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Errorf("error close stream: %s", err)
		return NewError(err, ErrorTypeSaveToStorage)
	}
	log.Infof("Route summary: %v", reply.Size)
	return nil
}

func (ec *ExtConnection) RestoreInfo(ctx context.Context, session string, w io.Writer, fileInfo *entity.FileInfo) error {
	conn, err := ec.createConnection()
	if err != nil {
		log.Errorf("error create grpc connection: %s", err)
		return err
	}
	defer conn.Close()
	c := pb.NewSecureStorageClient(conn)
	req := &pb.RestoreFileRequest{
		Session: session,
	}
	stream, err := c.RestoreFile(ctx, req)
	if err != nil {
		log.Errorf("error restore file: %s", err)
		return NewError(err, ErrorTypeRestoreFromStorage)
	}
	for {
		part, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Errorf("%v.ListFeatures(_) = _, %v", c, err)
			return NewError(err, ErrorTypeRestoreFromStorage)
		}
		_, err = w.Write(part.Content)
		if err != nil {
			log.Errorf("error write to file: %s", err)
			return NewError(err, ErrorTypeRestoreFromStorage)
		}
	}
	return nil
}
