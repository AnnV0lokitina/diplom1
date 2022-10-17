package handler

import (
	"context"
	"fmt"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
)

func (h *Handler) printCredentials(c entity.Credentials) {
	fmt.Printf(
		"Login: %s \nPassword: %s \nMeta: %s \n",
		c.Login,
		c.Password,
		c.Meta,
	)
}

func (h *Handler) printBankCard(c entity.BankCard) {
	fmt.Printf(
		"Number: %s \nExpDate: %s \nCardholder: %s \nCode: %s \nMeta: %s \n",
		c.Number,
		c.ExpDate,
		c.Cardholder,
		c.Code,
		c.Meta,
	)
}

func (h *Handler) printFile(f entity.File) {
	fmt.Printf(
		"Name: %s \nMeta: %s \n",
		f.Name,
		f.Meta,
	)
}

func (h *Handler) ShowCredentialsList(ctx context.Context) {
	credentials := h.service.ShowCredentialsList(ctx)
	for _, c := range credentials {
		h.printCredentials(c)
	}
}

func (h *Handler) ShowTextFilesList(ctx context.Context) {
	files := h.service.ShowTextFilesList(ctx)
	for _, f := range files {
		h.printFile(f)
	}
}

func (h *Handler) ShowBinaryDataList(ctx context.Context) {
	files := h.service.ShowBinaryDataList(ctx)
	for _, f := range files {
		h.printFile(f)
	}
}

func (h *Handler) ShowBankCardList(ctx context.Context) {
	credentials := h.service.ShowBankCardList(ctx)
	for _, c := range credentials {
		h.printBankCard(c)
	}
}

func (h *Handler) GetCredentialsByLogin(ctx context.Context, login string) {
	c := h.service.GetCredentialsByLogin(ctx, login)
	h.printCredentials(*c)
}

func (h *Handler) GetBankCardByNumber(ctx context.Context, number string) {
	c := h.service.GetBankCardByNumber(ctx, number)
	h.printBankCard(*c)
}

func (h *Handler) GetTextFileByName(ctx context.Context, name string, path string) error {
	//f := h.repo.GetTextFileByName(name)
	//f.Name
	return nil
}

func (h *Handler) GetBinaryDataByName(ctx context.Context, name string, path string) error {
	//f := h.repo.GetTextFileByName(name)
	//f.Name
	return nil
}
