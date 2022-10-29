package _interface

import (
	"context"
)

type Remote interface {
	SendInfo(ctx context.Context) error
	ReceiveInfo(ctx context.Context) error
	Register(ctx context.Context, login string, password string) error
	Login(ctx context.Context, login string, password string) error
}
