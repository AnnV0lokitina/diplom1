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
	session, err := s.session.Get()
	if err != nil {
		session = ""
	}
	err = s.receiveInfo(ctx, session)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	return s.repo.GetCredentialsList()
}

func (s *Service) ShowTextFilesList(ctx context.Context) []entity.File {
	session, err := s.session.Get()
	if err != nil {
		session = ""
	}
	err = s.receiveInfo(ctx, session)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	return s.repo.GetTextFileList()
}

func (s *Service) ShowBinaryDataList(ctx context.Context) []entity.File {
	session, err := s.session.Get()
	if err != nil {
		session = ""
	}
	err = s.receiveInfo(ctx, session)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	return s.repo.GetBinaryFileList()
}

func (s *Service) ShowBankCardList(ctx context.Context) []entity.BankCard {
	session, err := s.session.Get()
	if err != nil {
		session = ""
	}
	err = s.receiveInfo(ctx, session)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	return s.repo.GetBankCardList()
}

func (s *Service) GetCredentialsByLogin(ctx context.Context, login string) *entity.Credentials {
	session, err := s.session.Get()
	if err != nil {
		session = ""
	}
	err = s.receiveInfo(ctx, session)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	return s.repo.GetCredentialsByLogin(login)
}

func (s *Service) GetBankCardByNumber(ctx context.Context, number string) *entity.BankCard {
	session, err := s.session.Get()
	if err != nil {
		session = ""
	}
	err = s.receiveInfo(ctx, session)
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
	session, err := s.session.Get()
	if err != nil {
		session = ""
	}
	err = s.receiveInfo(ctx, session)
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
	session, err := s.session.Get()
	if err != nil {
		session = ""
	}
	err = s.receiveInfo(ctx, session)
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
