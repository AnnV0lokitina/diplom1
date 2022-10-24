package repo

import (
	"bufio"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
)

// AddTextFile Save text file information to storage.
func (r *Repo) AddTextFile(file entity.File, reader *bufio.Reader) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	err := r.enclosure.Save(file.Name, reader)
	if err != nil {
		return err
	}
	r.record.TextFileList = append(r.record.TextFileList, file)
	return r.writer.WriteRecord(r.record)
}

// AddBinaryFile Save binary file information to storage.
func (r *Repo) AddBinaryFile(file entity.File, reader *bufio.Reader) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	err := r.enclosure.Save(file.Name, reader)
	if err != nil {
		return err
	}
	r.record.BinaryFileList = append(r.record.BinaryFileList, file)
	return r.writer.WriteRecord(r.record)
}

// AddCredentials Save credentials information to storage.
func (r *Repo) AddCredentials(cred entity.Credentials) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.record.CredentialsList = append(r.record.CredentialsList, cred)
	return r.writer.WriteRecord(r.record)
}

// AddBankCard Save bank card information to storage.
func (r *Repo) AddBankCard(card entity.BankCard) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.record.BankCardList = append(r.record.BankCardList, card)
	return r.writer.WriteRecord(r.record)
}
