package handler

import "context"

func (h *Handler) Register(ctx context.Context, login string, password string) error {
	return h.service.Register(ctx, login, password)
}

func (h *Handler) Login(ctx context.Context, login string, password string) error {
	return h.service.Login(ctx, login, password)
}
