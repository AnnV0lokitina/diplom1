package handler

import (
	"context"
	"fmt"
)

// Register a new user.
func (h *Handler) Register(ctx context.Context, login string, password string) error {
	err := h.service.Register(ctx, login, password)
	if err != nil {
		return err
	}
	fmt.Println("Registration completed successfully")
	return nil
}

// Login Authorizes an existing user.
func (h *Handler) Login(ctx context.Context, login string, password string) error {
	err := h.service.Login(ctx, login, password)
	if err != nil {
		return err
	}
	fmt.Println("Authorization was successful")
	return nil
}
