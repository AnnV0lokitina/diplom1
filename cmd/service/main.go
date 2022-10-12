package main

import (
	"context"
	"fmt"
	dbPkg "github.com/AnnV0lokitina/diplom1/internal/db"
	handlerGRPCPkg "github.com/AnnV0lokitina/diplom1/internal/handler"
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
		signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

		<-c
		fmt.Println("cancel()")
		cancel()
	}()
	db, err := dbPkg.NewDB(ctx, cfg.DataBaseURI)
	if err != nil {
		log.WithError(err).Fatal("error db repo init")
	}
	defer db.Close(ctx)

	service := servicePkg.NewService(db)
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
