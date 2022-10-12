package service

func (s *Service) RemoveCredentialsByLogin(login string) error {
	return s.repo.RemoveCredentialsByLogin(login)
}

func (s *Service) RemoveTextByName(name string) error {
	return s.repo.RemoveTextFileByName(name)
}

func (s *Service) RemoveBinaryDataByName(name string) error {
	return s.repo.RemoveBinaryFileByName(name)
}

func (s *Service) RemoveBankCardByNumber(number string) error {
	return s.repo.RemoveBankCardByNumber(number)
}
