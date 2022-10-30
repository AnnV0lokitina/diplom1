package service

import (
	"context"
	log "github.com/sirupsen/logrus"
)

func (s *Service) RemoveCredentialsByLogin(ctx context.Context, login string) error {
	err := s.r.ReceiveInfo(ctx)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	err = s.repo.RemoveCredentialsByLogin(login)
	if err != nil {
		return err
	}
	return s.r.SendInfo(ctx)
}

func (s *Service) RemoveTextByName(ctx context.Context, name string) error {
	err := s.r.ReceiveInfo(ctx)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	err = s.repo.RemoveTextFileByName(name)
	if err != nil {
		return err
	}
	return s.r.SendInfo(ctx)
}

func (s *Service) RemoveBinaryDataByName(ctx context.Context, name string) error {
	err := s.r.ReceiveInfo(ctx)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	err = s.repo.RemoveBinaryFileByName(name)
	if err != nil {
		return err
	}
	return s.r.SendInfo(ctx)
}

func (s *Service) RemoveBankCardByNumber(ctx context.Context, number string) error {
	err := s.r.ReceiveInfo(ctx)
	if err != nil {
		log.Info("receive info: " + err.Error())
	}
	err = s.repo.RemoveBankCardByNumber(number)
	if err != nil {
		return err
	}
	return s.r.SendInfo(ctx)
}
