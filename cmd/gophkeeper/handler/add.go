package handler

import "context"

func (h *Handler) AddCredentials(ctx context.Context, login string, password string, meta string) error {
	return h.service.AddCredentials(ctx, login, password, meta)
}

func (h *Handler) AddTextFromFile(path string, meta string) error {
	return h.service.AddTextFromFile(path, meta)
}

func (h *Handler) AddBinaryDataFromFile(path string, meta string) error {
	return h.service.AddBinaryDataFromFile(path, meta)
}

func (h *Handler) AddBankCard(number string, exp string, cardholder string, code string, meta string) error {
	return h.AddBankCard(number, exp, cardholder, code, meta)
}
