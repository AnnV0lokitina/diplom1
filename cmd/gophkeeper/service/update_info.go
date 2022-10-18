package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"io"
)

// sendInfo Send information to server.
func (s *Service) sendInfo(ctx context.Context, session string) error {
	log.Info("send info: read file start")
	err := s.repo.CreateZIP()
	if err != nil {
		log.Error("send info: create zip error")
		return err
	}
	fileInfo, err := s.repo.GetInfo()
	if err != nil {
		log.Error("send info: no file to send")
		return err
	}
	r, w := io.Pipe()
	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		defer w.Close()
		return s.repo.ReadFileByChunks(w)
	})
	g.Go(func() error {
		defer r.Close()
		return s.connection.StoreInfo(ctx, session, r, fileInfo)
	})
	if err := g.Wait(); err != nil {
		return err
	}
	log.Info("send info: read file end")
	return nil
}

// receiveInfo Receive information from server.
func (s *Service) receiveInfo(ctx context.Context, session string) error {
	log.Info("receive info: read file start")
	fileInfo, err := s.repo.GetInfo()
	if err != nil {
		log.Info("has no stored info to update, receive info")
	}
	r, w := io.Pipe()
	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		defer r.Close()
		return s.repo.WriteFileByChunks(r)
	})
	g.Go(func() error {
		defer w.Close()
		return s.connection.RestoreInfo(ctx, session, w, fileInfo)
	})
	if err != nil {
		return err
	}
	if err := g.Wait(); err != nil {
		return err
	}
	err = s.repo.UnpackZIP()
	if err != nil {
		log.Error("send info: unpack zip error")
		return err
	}
	log.Info("receive info: read file end")
	return nil
}
