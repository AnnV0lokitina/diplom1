package handler

import (
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

func (h *Handler) ShowCredentialsList() {
	credentials := h.service.ShowCredentialsList()
	for _, c := range credentials {
		h.printCredentials(c)
	}
}

func (h *Handler) ShowTextFilesList() {
	files := h.service.ShowTextFilesList()
	for _, f := range files {
		h.printFile(f)
	}
}

func (h *Handler) ShowBinaryDataList() {
	files := h.service.ShowBinaryDataList()
	for _, f := range files {
		h.printFile(f)
	}
}

func (h *Handler) ShowBankCardList() {
	credentials := h.service.ShowBankCardList()
	for _, c := range credentials {
		h.printBankCard(c)
	}
}

func (h *Handler) GetCredentialsByLogin(login string) {
	c := h.service.GetCredentialsByLogin(login)
	h.printCredentials(*c)
}

func (h *Handler) GetBankCardByNumber(number string) {
	c := h.service.GetBankCardByNumber(number)
	h.printBankCard(*c)
}

func (h *Handler) GetTextFileByName(name string, path string) error {
	//f := h.repo.GetTextFileByName(name)
	//f.Name
	return nil
}

func (h *Handler) GetBinaryDataByName(name string, path string) error {
	//f := h.repo.GetTextFileByName(name)
	//f.Name
	return nil
}
