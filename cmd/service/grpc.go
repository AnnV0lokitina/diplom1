package main

import (
	"github.com/AnnV0lokitina/diplom1/internal/handler"
	"google.golang.org/grpc"
)

type GRPCService struct {
	Handler *handler.Handler
	Server  *grpc.Server
}
