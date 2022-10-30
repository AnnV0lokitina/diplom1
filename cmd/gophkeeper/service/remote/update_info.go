package remote

import (
	"context"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/service/interface"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"io"
)

type Remote struct {
	Repo       _interface.RepoZip
	Connection _interface.ExtConnection
	Session    _interface.Session
}

// SendInfo Send information to server.
func (rs *Remote) SendInfo(ctx context.Context) error {
	session, err := rs.Session.Get()
	if err != nil {
		session = ""
	}
	log.Info("send info: read file start")
	err = rs.Repo.CreateZIP()
	if err != nil {
		log.Error("send info: create zip error")
		return err
	}
	fileInfo, err := rs.Repo.GetInfo()
	if err != nil {
		log.Error("send info: no file to send")
		return err
	}
	r, w := io.Pipe()
	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		defer w.Close()
		return rs.Repo.ReadFileByChunks(w)
	})
	g.Go(func() error {
		defer r.Close()
		return rs.Connection.StoreInfo(ctx, session, r, fileInfo)
	})
	if err := g.Wait(); err != nil {
		return err
	}
	log.Info("send info: read file end")
	return nil
}

// ReceiveInfo Receive information from server.
func (rs *Remote) ReceiveInfo(ctx context.Context) error {
	session, err := rs.Session.Get()
	if err != nil {
		session = ""
	}
	log.Info("receive info: read file start")
	fileInfo, err := rs.Repo.GetInfo()
	if err != nil {
		log.Info("has no stored info to update, receive info")
	}
	r, w := io.Pipe()
	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		defer r.Close()
		return rs.Repo.WriteFileByChunks(r)
	})
	g.Go(func() error {
		defer w.Close()
		return rs.Connection.RestoreInfo(ctx, session, w, fileInfo)
	})
	//if err != nil {
	//	return err
	//}
	if err := g.Wait(); err != nil {
		return err
	}
	err = rs.Repo.UnpackZIP()
	if err != nil {
		log.Error("send info: unpack zip error")
		return err
	}
	log.Info("receive info: read file end")
	return nil
}
