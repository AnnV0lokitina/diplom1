package service

import (
	"context"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	log "github.com/sirupsen/logrus"
)

func (s *Service) ShowCredentialsList(ctx context.Context) []entity.Credentials {
	session, err := GetSession()
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
	session, err := GetSession()
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
	session, err := GetSession()
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
	session, err := GetSession()
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
	session, err := GetSession()
	if err != nil {
		session = ""
	}
	err = s.receiveInfo(ctx, session)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	return s.repo.GetCredentialsByNumber(login)
}

func (s *Service) GetBankCardByNumber(ctx context.Context, number string) *entity.BankCard {
	session, err := GetSession()
	if err != nil {
		session = ""
	}
	err = s.receiveInfo(ctx, session)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	return s.repo.GetBankCardByNumber(number)
}

func (s *Service) GetTextFileByName(ctx context.Context, name string, path string) error {
	//f := h.repo.GetTextFileByName(name)
	//f.Name
	return nil
}

func (s *Service) GetBinaryDataByName(ctx context.Context, name string, path string) error {
	//f := h.repo.GetTextFileByName(name)
	//f.Name
	return nil
}
