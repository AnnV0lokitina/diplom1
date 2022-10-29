package service

import (
	"bufio"
	"context"
	"errors"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	log "github.com/sirupsen/logrus"
	"os"
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

// AddTextFromFile Saves a text from file to storage.
func (s *Service) AddTextFromFile(ctx context.Context, path string, meta string) error {
	err := s.r.ReceiveInfo(ctx)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	stat, err := os.Stat(path)
	if os.IsNotExist(err) || stat.Size() == 0 {
		return errors.New("no source file")
	}
	file, err := os.OpenFile(path, os.O_RDONLY, 0777)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(file)
	if err != nil {
		return err
	}
	info := entity.File{
		Name: stat.Name(),
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
	stat, err := os.Stat(path)
	if os.IsNotExist(err) || stat.Size() == 0 {
		return errors.New("no source file")
	}
	file, err := os.OpenFile(path, os.O_RDONLY, 0777)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(file)
	if err != nil {
		return err
	}
	info := entity.File{
		Name: stat.Name(),
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
