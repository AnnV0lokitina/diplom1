package handler

import (
	"context"
	"fmt"
)

// RemoveCredentialsByLogin Removes a pair of login password.
func (h *Handler) RemoveCredentialsByLogin(ctx context.Context, login string) error {
	err := h.service.RemoveCredentialsByLogin(ctx, login)
	if err != nil {
		return err
	}
	fmt.Println("Credentials successfully removed")
	return nil
}

// RemoveTextByName Delete text by title.
func (h *Handler) RemoveTextByName(ctx context.Context, name string) error {
	err := h.service.RemoveTextByName(ctx, name)
	if err != nil {
		return err
	}
	fmt.Println("Text successfully removed")
	return nil
}

// RemoveBinaryDataByName Delete binary file by title.
func (h *Handler) RemoveBinaryDataByName(ctx context.Context, name string) error {
	err := h.service.RemoveBinaryDataByName(ctx, name)
	if err != nil {
		return err
	}
	fmt.Println("Binary files successfully removed")
	return nil
}

// RemoveBankCardByNumber Deletes a bank card by number.
func (h *Handler) RemoveBankCardByNumber(ctx context.Context, number string) error {
	err := h.service.RemoveBankCardByNumber(ctx, number)
	if err != nil {
		return err
	}
	fmt.Println("Bank card successfully removed")
	return nil
}
