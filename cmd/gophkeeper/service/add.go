package service

import (
	"bufio"
	"context"
	"errors"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	"golang.org/x/sync/errgroup"
	"io"
	"os"
)

func (s *Service) sendInfo(ctx context.Context, session string) error {
	r, w := io.Pipe()
	fileInfo, err := s.repo.GetInfo()
	if err != nil {
		return err
	}
	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		defer w.Close()
		return s.repo.ReadFileByChunks(w)
	})
	err = s.connection.StoreInfo(ctx, session, r, fileInfo)
	if err != nil {
		return err
	}
	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

func (s *Service) receiveInfo(ctx context.Context, session string) error {
	r, w := io.Pipe()
	fileInfo, err := s.repo.GetInfo()
	if err != nil {
		// если пустой обновить независимо от даты
		return err
	}
	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		defer r.Close()
		return s.repo.WriteFileByChunks(r)
	})
	err = s.connection.RestoreInfo(ctx, session, w, fileInfo)
	if err != nil {
		return err
	}
	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

// 1kb
func (s *Service) AddCredentials(ctx context.Context, login string, password string, meta string) error {
	session := os.Getenv("EXT_SESSION")
	err := s.receiveInfo(ctx, session)
	if err != nil {
		return err
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
	return s.sendInfo(ctx, session)
}

func (s *Service) AddTextFromFile(path string, meta string) error {
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
	return s.repo.AddTextFile(info, reader)
}

func (s *Service) AddBinaryDataFromFile(path string, meta string) error {
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
	return s.repo.AddBinaryFile(info, reader)
}

func (s *Service) AddBankCard(number string, exp string, cardholder string, code string, meta string) error {
	card := entity.BankCard{
		Number:     number,
		ExpDate:    exp,
		Cardholder: cardholder,
		Code:       code,
		Meta:       meta,
	}
	return s.repo.AddBankCard(card)
}
