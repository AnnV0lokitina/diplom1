package service

import (
	"context"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	log "github.com/sirupsen/logrus"
)

// AddCredentials Saves a pair of login and password.
func (s *Service) AddCredentials(ctx context.Context, login string, password string, meta string) error {
	err := s.r.ReceiveInfo(ctx)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	cred := entity.Credentials{
		Login:    login,
		Password: password,
		Meta:     meta,
	}
	err = s.repo.AddCredentials(cred)
	if err != nil {
		return err
	}
	return s.r.SendInfo(ctx)
}

// AddText Saves a text to storage.
func (s *Service) AddText(ctx context.Context, text string, name string, meta string) error {
	err := s.r.ReceiveInfo(ctx)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	reader := entity.NewTextReadCloser(text)
	info := entity.File{
		Name: name,
		Meta: meta,
	}
	err = s.repo.AddTextFile(info, reader)
	if err != nil {
		return err
	}
	return s.r.SendInfo(ctx)
}

// AddBinaryDataFromFile Saves a binary file to storage.
func (s *Service) AddBinaryDataFromFile(ctx context.Context, path string, meta string) error {
	err := s.r.ReceiveInfo(ctx)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	name, reader, err := s.ext.Open(path)
	info := entity.File{
		Name: name,
		Meta: meta,
	}
	err = s.repo.AddBinaryFile(info, reader)
	if err != nil {
		return err
	}
	return s.r.SendInfo(ctx)
}

// AddBankCard Saves a bank card to storage.
func (s *Service) AddBankCard(
	ctx context.Context,
	number string,
	exp string,
	cardholder string,
	code string,
	meta string,
) error {
	err := s.r.ReceiveInfo(ctx)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	card := entity.BankCard{
		Number:     number,
		ExpDate:    exp,
		Cardholder: cardholder,
		Code:       code,
		Meta:       meta,
	}
	err = s.repo.AddBankCard(card)
	if err != nil {
		return err
	}
	return s.r.SendInfo(ctx)
}
