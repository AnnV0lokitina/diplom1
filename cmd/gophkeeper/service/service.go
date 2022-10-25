package service

type Service struct {
	repo       Repo
	connection ExtConnection
	session    Session
}

func NewService(repo Repo, conn ExtConnection, s Session) *Service {
	return &Service{
		repo:       repo,
		connection: conn,
		session:    s,
	}
}
