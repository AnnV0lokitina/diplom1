package repo

import (
	"errors"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	"io"
)

var errorDuplicate = errors.New("duplicate")

// AddTextFile Save text file information to storage.
func (r *Repo) AddTextFile(file entity.File, reader io.Reader) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	oldFileList := make([]entity.File, 0, len(r.record.TextFileList))
	for _, item := range r.record.TextFileList {
		if item.Name == file.Name {
			return errorDuplicate
		}
		oldFileList = append(oldFileList, item)
	}
	r.record.TextFileList = append(r.record.TextFileList, file)
	err := r.enclosure.Save(file.Name, reader)
	if err != nil {
		r.record.TextFileList = oldFileList
		return err
	}
	err = r.writer.WriteRecord(r.record)
	if err != nil {
		r.record.TextFileList = oldFileList
		return err
	}
	return nil
}

// AddBinaryFile Save binary file information to storage.
func (r *Repo) AddBinaryFile(file entity.File, reader io.Reader) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	oldFileList := make([]entity.File, 0, len(r.record.BinaryFileList))
	for _, item := range r.record.BinaryFileList {
		if item.Name == file.Name {
			return errorDuplicate
		}
		oldFileList = append(oldFileList, item)
	}
	r.record.BinaryFileList = append(r.record.BinaryFileList, file)
	err := r.enclosure.Save(file.Name, reader)
	if err != nil {
		r.record.BinaryFileList = oldFileList
		return err
	}
	err = r.writer.WriteRecord(r.record)
	if err != nil {
		r.record.BinaryFileList = oldFileList
		return err
	}
	return nil
}

// AddCredentials Save credentials information to storage.
func (r *Repo) AddCredentials(cred entity.Credentials) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	oldList := make([]entity.Credentials, 0, len(r.record.CredentialsList))
	for _, item := range r.record.CredentialsList {
		if item.Login == cred.Login {
			return errorDuplicate
		}
		oldList = append(oldList, item)
	}
	r.record.CredentialsList = append(r.record.CredentialsList, cred)
	err := r.writer.WriteRecord(r.record)
	if err != nil {
		r.record.CredentialsList = oldList
		return err
	}
	return nil
}

// AddBankCard Save bank card information to storage.
func (r *Repo) AddBankCard(card entity.BankCard) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	oldList := make([]entity.BankCard, 0, len(r.record.BankCardList))
	for _, item := range r.record.BankCardList {
		if item.Number == card.Number {
			return errorDuplicate
		}
		oldList = append(oldList, item)
	}
	r.record.BankCardList = append(r.record.BankCardList, card)
	err := r.writer.WriteRecord(r.record)
	if err != nil {
		r.record.BankCardList = oldList
		return err
	}
	return nil
}
