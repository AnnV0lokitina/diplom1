package main

import (
	"context"
	pb "github.com/AnnV0lokitina/diplom1/proto"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
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

	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		log.Fatal(err)
	}

	pb.RegisterSecureStorageServer(app.g.Server, app.g.Handler)
	log.Println("gRPC server starts")
	if err := app.g.Server.Serve(listen); err != nil {
		log.Fatal(err)
	}

	go func() {
		<-ctx.Done()

		app.g.Server.GracefulStop()
		log.Println("stop grpc")
		httpShutdownCh <- struct{}{}
	}()

	<-httpShutdownCh
	if err == http.ErrServerClosed {
		return nil
	}
	return err
}
