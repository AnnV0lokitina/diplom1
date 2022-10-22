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
	if c == nil {
		fmt.Println("Login password pair not found")
		return
	}
	h.printCredentials(*c)
}

func (h *Handler) GetBankCardByNumber(ctx context.Context, number string) {
	c := h.service.GetBankCardByNumber(ctx, number)
	if c == nil {
		fmt.Println("Bank card not found")
		return
	}
	h.printBankCard(*c)
}

func (h *Handler) GetTextFileByName(ctx context.Context, name string, path string) error {
	f, err := h.service.UploadTextFileByNameIntoPath(ctx, name, path)
	if err != nil {
		fmt.Println("Error while uploading text file")
		return err
	}
	fmt.Println(fmt.Sprintf("Successfully upload text file %s", f.Name))
	return nil
}

func (h *Handler) GetBinaryDataByName(ctx context.Context, name string, path string) error {
	f, err := h.service.UploadBinaryFileByNameIntoPath(ctx, name, path)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error while uploading binary file %s", f.Name))
		return nil
	}
	fmt.Println("Successfully upload binary file")
	return nil
}
