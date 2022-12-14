package main

import (
	"context"
	"fmt"
	pb "github.com/AnnV0lokitina/diplom1/proto"
	log "github.com/sirupsen/logrus"
	"net"
)

type App struct {
	g *GRPCService
}

func NewApp(grpcService *GRPCService) *App {
	return &App{
		g: grpcService,
	}
}

func (app *App) Run(ctx context.Context, serverAddress string) error {
	httpShutdownCh := make(chan struct{})

	go func() {
		<-ctx.Done()
		fmt.Println("done")

		app.g.Server.GracefulStop()
		log.Println("stop grpc")
		httpShutdownCh <- struct{}{}
	}()

	listen, err := net.Listen("tcp", serverAddress)
	if err != nil {
		return err
	}

	pb.RegisterSecureStorageServer(app.g.Server, app.g.Handler)
	log.Println("gRPC server starts")
	if err := app.g.Server.Serve(listen); err != nil {
		return err
	}

	<-httpShutdownCh
	return nil
}
