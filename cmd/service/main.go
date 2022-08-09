package main

import (
	"context"
	handlerGRPCPkg "github.com/AnnV0lokitina/diplom1/internal/handler"
	repoPkg "github.com/AnnV0lokitina/diplom1/internal/repo"
	servicePkg "github.com/AnnV0lokitina/diplom1/internal/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := initConfig()
	initParams(cfg)
	err := doMigrates(cfg.DataBaseURI)
	if err != nil {
		log.WithError(err).Fatal("migrations error")
	}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		cancel()
	}()
	repo, err := repoPkg.NewRepo(ctx, cfg.DataBaseURI)
	if err != nil {
		log.WithError(err).Fatal("error repo init")
	}
	defer repo.Close(ctx)

	service := servicePkg.NewService(repo)
	handler := &GRPCService{
		Handler: handlerGRPCPkg.NewHandler(service),
		Server:  grpc.NewServer(),
	}
	application := NewApp(handler)

	err = application.Run(ctx, cfg.RunAddress)
	if err != nil {
		log.Fatal(err)
	}
}