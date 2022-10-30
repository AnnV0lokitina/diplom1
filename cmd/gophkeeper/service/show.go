package service

import (
	"bytes"
	"context"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	log "github.com/sirupsen/logrus"
	"io"
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

func (s *Service) GetTextByName(ctx context.Context, name string) (*entity.File, string, error) {
	err := s.r.ReceiveInfo(ctx)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	f, reader, err := s.repo.GetTextFileByName(name)
	if err != nil {
		return nil, "", err
	}
	defer reader.Close()
	textBuf := bytes.Buffer{}
	_, err = io.Copy(&textBuf, reader)
	if err != nil {
		return nil, "", err
	}
	return f, textBuf.String(), nil
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
	err = s.ext.Save(outFilePath, reader)
	if err != nil {
		return nil, err
	}
	return f, nil
}
