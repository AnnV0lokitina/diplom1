package remote

import (
	"context"
)

// Register a user on a remote server.
func (rs *Remote) Register(ctx context.Context, login string, password string) error {
	session, err := rs.Connection.Register(ctx, login, password)
	if err != nil {
		return err
	}
	return rs.Session.Save(session)
}

// Login Authorizes a user on a remote server.
func (rs *Remote) Login(ctx context.Context, login string, password string) error {
	session, err := rs.Connection.Login(ctx, login, password)
	if err != nil {
		return err
	}
	return rs.Session.Save(session)
}
