package handler

import (
	"context"
	"fmt"
)

// AddCredentials Processes a request to add a login/password pair.
func (h *Handler) AddCredentials(ctx context.Context, login string, password string, meta string) error {
	err := h.service.AddCredentials(ctx, login, password, meta)
	if err != nil {
		return err
	}
	fmt.Println("Login password pair added successfully")
	return nil
}

// AddTextFromFile Handles a request to add text from a file.
func (h *Handler) AddTextFromFile(ctx context.Context, text string, name string, meta string) error {
	err := h.service.AddText(ctx, text, name, meta)
	if err != nil {
		return err
	}
	fmt.Println("Text from file added successfully")
	return nil
}

// AddBinaryDataFromFile Handles a request to add a binary file.
func (h *Handler) AddBinaryDataFromFile(ctx context.Context, path string, meta string) error {
	err := h.service.AddBinaryDataFromFile(ctx, path, meta)
	if err != nil {
		return err
	}
	fmt.Println("Binary added successfully")
	return nil
}

// AddBankCard Processes a request to add a bank card.
func (h *Handler) AddBankCard(
	ctx context.Context,
	number string,
	exp string,
	cardholder string,
	code string,
	meta string,
) error {
	err := h.AddBankCard(ctx, number, exp, cardholder, code, meta)
	if err != nil {
		return err
	}
	fmt.Println("Bank card added successfully")
	return nil
}
