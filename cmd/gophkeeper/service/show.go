package service

import "github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"

func (s *Service) ShowCredentialsList() []entity.Credentials {
	return s.repo.GetCredentialsList()
}

func (s *Service) ShowTextFilesList() []entity.File {
	return s.repo.GetTextFileList()
}

func (s *Service) ShowBinaryDataList() []entity.File {
	return s.repo.GetBinaryFileList()
}

func (s *Service) ShowBankCardList() []entity.BankCard {
	return s.repo.GetBankCardList()
}

func (s *Service) GetCredentialsByLogin(login string) *entity.Credentials {
	return s.repo.GetCredentialsByNumber(login)
}

func (s *Service) GetBankCardByNumber(number string) *entity.BankCard {
	return s.repo.GetBankCardByNumber(number)
}

func (s *Service) GetTextFileByName(name string, path string) error {
	//f := h.repo.GetTextFileByName(name)
	//f.Name
	return nil
}

func (s *Service) GetBinaryDataByName(name string, path string) error {
	//f := h.repo.GetTextFileByName(name)
	//f.Name
	return nil
}
