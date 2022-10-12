package external

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log"
	"os"
	// ...
	pb "github.com/AnnV0lokitina/diplom1/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ExtConnection struct {
	//c pb.SecureStorageClient
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
		return nil, err
	}
	return conn, nil
}

func (ec *ExtConnection) Register(ctx context.Context, login string, password string) (string, error) {
	conn, err := ec.createConnection()
	if err != nil {
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
		return "", err
	}
	return resp.GetSession(), nil
}

func (ec *ExtConnection) Login(ctx context.Context, login string, password string) (string, error) {
	conn, err := ec.createConnection()
	if err != nil {
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
		return "", err
	}
	return resp.GetSession(), nil
}

func (ec *ExtConnection) StoreInfo(ctx context.Context, session string, reader io.Reader, info os.FileInfo) error {
	conn, err := ec.createConnection()
	if err != nil {
		return err
	}
	defer conn.Close()
	c := pb.NewSecureStorageClient(conn)
	stream, err := c.StoreFile(ctx)
	if err != nil {
		return fmt.Errorf("%v.StoreFile(_) = _, %v", c, err)
	}
	buf := make([]byte, 16)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		part := pb.StoreFileRequest{
			Info: &pb.FileInfo{
				Name: info.Name(),
				Type: "txt",
				Time: timestamppb.New(info.ModTime()),
			},
			Content: buf[:n],
			Session: session,
		}
		fmt.Println("send", string(buf[:n]))
		if err := stream.Send(&part); err != nil {
			fmt.Println(err)
			return fmt.Errorf("%v.Send(%v) = %v", stream, part, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Route summary: %v", reply.Size)
	return nil
}

func (ec *ExtConnection) RestoreInfo(ctx context.Context, session string, w io.Writer, fileInfo os.FileInfo) error {
	conn, err := ec.createConnection()
	if err != nil {
		return err
	}
	defer conn.Close()
	c := pb.NewSecureStorageClient(conn)
	req := &pb.RestoreFileRequest{
		Session: session,
	}
	stream, err := c.RestoreFile(ctx, req)
	if err != nil {
		return err
	}
	for {
		part, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", c, err)
		}
		_, err = w.Write(part.Content)
		if err != nil {
			return err
		}
		//log.Println(part)
	}
	return nil
}
