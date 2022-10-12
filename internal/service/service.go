package service

import (
	"context"
	"errors"
	"github.com/AnnV0lokitina/diplom1/internal/entity"
	labelError "github.com/AnnV0lokitina/diplom1/pkg/error"
	"github.com/AnnV0lokitina/diplom1/pkg/file"
	log "github.com/sirupsen/logrus"
	"io"
)

type DB interface {
	Close(ctx context.Context) error
	CreateUser(
		ctx context.Context,
		sessionID string,
		login string,
		passwordHash string,
	) error
	AuthUser(
		ctx context.Context,
		login string,
		passwordHash string,
	) (int, error)
	AddUserSession(ctx context.Context, user *entity.User) error
	GetUserBySessionID(ctx context.Context, activeSessionID string) (*entity.User, error)
}

type Service struct {
	db DB
}

func NewService(db DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) RegisterUser(ctx context.Context, login string, password string) (*entity.User, error) {
	passwordHash := entity.CreatePasswordHash(password)
	sessionID, err := entity.GenerateSessionID()
	if err != nil {
		return nil, err
	}
	err = s.db.CreateUser(ctx, sessionID, login, passwordHash)
	if err != nil {
		return nil, err
	}
	user := &entity.User{
		ActiveSessionID: sessionID,
		Login:           login,
	}
	return user, nil
}

func (s *Service) LoginUser(ctx context.Context, login string, password string) (*entity.User, error) {
	passwordHash := entity.CreatePasswordHash(password)
	sessionID, err := entity.GenerateSessionID()
	if err != nil {
		return nil, err
	}
	userID, err := s.db.AuthUser(ctx, login, passwordHash)
	if err != nil {
		var labelErr *labelError.LabelError
		if errors.As(err, &labelErr) && labelErr.Label == labelError.TypeNotFound {
			log.Info("user not found")
			return nil, labelError.NewLabelError(labelError.TypeUnauthorized, errors.New("user unauthorized"))
		}
		return nil, err
	}
	user := &entity.User{
		ID:              userID,
		Login:           login,
		ActiveSessionID: sessionID,
	}
	err = s.db.AddUserSession(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) authorizeUser(ctx context.Context, sessionID string) (*entity.User, error) {
	user, err := s.db.GetUserBySessionID(ctx, sessionID)
	if err != nil {
		var labelErr *labelError.LabelError
		if errors.As(err, &labelErr) && labelErr.Label == labelError.TypeNotFound {
			log.Info("user not found")
			return nil, labelError.NewLabelError(labelError.TypeUnauthorized, errors.New("user unauthorized"))
		}
		return nil, err
	}
	return user, nil
}

func (s *Service) RestoreFile(session string, fileType string, fileName string, w io.Writer, time string) error {
	f := file.File{
		Path: "data.json",
	}
	return f.ReadByChunks(w)
}

func (s *Service) StoreFile(session string, fileType string, fileName string, r io.Reader, time string) error {
	f := file.File{
		Path: "data.json",
	}
	return f.WriteByChunks(r)
	//w, err := entity.NewWriter("data.json")
	//if err != nil {
	//	return err
	//}
	//defer w.Close()
	//b := make([]byte, 8)
	//for {
	//	log.Println("read start")
	//	n, err := r.Read(b)
	//	log.Printf("n = %v err = %v b = %v\n", n, err, string(b))
	//	log.Printf("b[:n] = %q\n", b[:n])
	//	if err == io.EOF || n == 0 {
	//		log.Println("eof")
	//		break
	//	}
	//	n, err = w.Write(b[:n])
	//	log.Println(n)
	//	if err != nil {
	//		return err
	//	}
	//}
	//return nil
}
