package service

import (
	"context"
	log "github.com/sirupsen/logrus"
)

func (s *Service) RemoveCredentialsByLogin(ctx context.Context, login string) error {
	session, err := GetSession()
	if err != nil {
		session = ""
	}
	err = s.receiveInfo(ctx, session)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	err = s.repo.RemoveCredentialsByLogin(login)
	if err != nil {
		return err
	}
	return s.sendInfo(ctx, session)
}

func (s *Service) RemoveTextByName(ctx context.Context, name string) error {
	session, err := GetSession()
	if err != nil {
		session = ""
	}
	err = s.receiveInfo(ctx, session)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	err = s.repo.RemoveTextFileByName(name)
	if err != nil {
		return err
	}
	return s.sendInfo(ctx, session)
}

func (s *Service) RemoveBinaryDataByName(ctx context.Context, name string) error {
	session, err := GetSession()
	if err != nil {
		session = ""
	}
	err = s.receiveInfo(ctx, session)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	err = s.repo.RemoveBinaryFileByName(name)
	if err != nil {
		return err
	}
	return s.sendInfo(ctx, session)
}

func (s *Service) RemoveBankCardByNumber(ctx context.Context, number string) error {
	session, err := GetSession()
	if err != nil {
		session = ""
	}
	err = s.receiveInfo(ctx, session)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	err = s.repo.RemoveBankCardByNumber(number)
	if err != nil {
		return err
	}
	return s.sendInfo(ctx, session)
}
