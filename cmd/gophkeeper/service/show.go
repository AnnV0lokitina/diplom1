package service

import (
	"context"
	"fmt"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func (s *Service) ShowCredentialsList(ctx context.Context) []entity.Credentials {
	err := s.r.ReceiveInfo(ctx)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	return s.repo.GetCredentialsList()
}

func (s *Service) ShowTextFilesList(ctx context.Context) []entity.File {
	err := s.r.ReceiveInfo(ctx)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	return s.repo.GetTextFileList()
}

func (s *Service) ShowBinaryDataList(ctx context.Context) []entity.File {
	err := s.r.ReceiveInfo(ctx)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	return s.repo.GetBinaryFileList()
}

func (s *Service) ShowBankCardList(ctx context.Context) []entity.BankCard {
	err := s.r.ReceiveInfo(ctx)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	return s.repo.GetBankCardList()
}

func (s *Service) GetCredentialsByLogin(ctx context.Context, login string) *entity.Credentials {
	err := s.r.ReceiveInfo(ctx)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	return s.repo.GetCredentialsByLogin(login)
}

func (s *Service) GetBankCardByNumber(ctx context.Context, number string) *entity.BankCard {
	err := s.r.ReceiveInfo(ctx)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	return s.repo.GetBankCardByNumber(number)
}

func (s *Service) UploadTextFileByNameIntoPath(
	ctx context.Context,
	name string,
	outFilePath string,
) (*entity.File, error) {
	err := s.r.ReceiveInfo(ctx)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	f, reader, err := s.repo.GetTextFileByName(name)
	if err != nil {
		return nil, err
	}
	fo, err := os.Create(outFilePath)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fo, reader)
	if err != nil {
		return nil, err
	}
	fmt.Println(f.Name)
	return f, nil
}

func (s *Service) UploadBinaryFileByNameIntoPath(
	ctx context.Context,
	name string,
	outFilePath string,
) (*entity.File, error) {
	err := s.r.ReceiveInfo(ctx)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	f, reader, err := s.repo.GetBinaryFileByName(name)
	if err != nil {
		return nil, err
	}
	fo, err := os.Create(outFilePath)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fo, reader)
	if err != nil {
		return nil, err
	}
	fmt.Println(f.Name)
	return f, nil
}
