package service

type Service struct {
	repo       Repo
	connection ExtConnection
}

func NewService(repo Repo, conn ExtConnection) *Service {
	return &Service{
		repo:       repo,
		connection: conn,
	}
}
