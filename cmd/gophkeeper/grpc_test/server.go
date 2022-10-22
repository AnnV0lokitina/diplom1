package grpctest

import (
	"context"
	"fmt"
	pb "github.com/AnnV0lokitina/diplom1/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type GRPCService struct {
	Handler *Handler
	Server  *grpc.Server
}

type Server struct {
	target     string
	cancel     *context.CancelFunc
	shutdownCh chan struct{}
}

func NewServer(target string) *Server {
	return &Server{
		target: target,
	}
}

func (ts *Server) Start(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

		<-c
		fmt.Println("cancel()")
		cancel()
	}()
	ts.cancel = &cancel

	handler := Handler{}
	server := grpc.NewServer()

	ts.shutdownCh = make(chan struct{})
	startCh := make(chan struct{})

	go func() {
		<-ctx.Done()
		fmt.Println("done")

		server.GracefulStop()
		log.Println("stop grpc test server")
		ts.shutdownCh <- struct{}{}
	}()

	go func() {
		listen, err := net.Listen("tcp", ts.target)
		if err != nil {
			log.Fatal(err)
		}

		pb.RegisterSecureStorageServer(server, &handler)
		log.Println("gRPC test server starts")
		startCh <- struct{}{}
		if err := server.Serve(listen); err != nil {
			log.Fatal(err)
		}
	}()
	<-startCh
}

func (ts *Server) Stop() {
	(*ts.cancel)()
	<-ts.shutdownCh
}
