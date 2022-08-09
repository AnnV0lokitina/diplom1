package repo

import "github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"

// AddTextFileName Save text file information to storage.
func (r *Repo) AddTextFileName(file entity.File) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.record.TextFileList = append(r.record.TextFileList, file)
	return r.writer.WriteRecord(r.record)
}

// AddBinaryFileName Save binary file information to storage.
func (r *Repo) AddBinaryFileName(file entity.File) error {
	r.mu.Lock()
	defer r.mu.Unlock()
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
