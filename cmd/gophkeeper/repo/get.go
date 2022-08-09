package repo

import "github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"

// GetTextFileList Get all text files information from storage.
func (r *Repo) GetTextFileList() []entity.File {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.record.TextFileList
}

// GetTextFileByName Get text files information from storage by name.
func (r *Repo) GetTextFileByName(name string) *entity.File {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i := range r.record.TextFileList {
		if r.record.TextFileList[i].Name == name {
			return &r.record.TextFileList[i]
		}
	}
	return nil
}

// GetBinaryFileList Get all binary files information to storage.
func (r *Repo) GetBinaryFileList() []entity.File {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.record.BinaryFileList
}

// GetBinaryFileByNumber Get binary file from storage by name.
func (r *Repo) GetBinaryFileByNumber(name string) *entity.File {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i := range r.record.BinaryFileList {
		if r.record.BinaryFileList[i].Name == name {
			return &r.record.BinaryFileList[i]
		}
	}
	return nil
}

// GetCredentialsList Get credentials from storage.
func (r *Repo) GetCredentialsList() []entity.Credentials {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.record.CredentialsList
}

// GetCredentialsByNumber Get credentials from storage by number.
func (r *Repo) GetCredentialsByNumber(login string) *entity.Credentials {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i := range r.record.BinaryFileList {
		if r.record.CredentialsList[i].Login == login {
			return &r.record.CredentialsList[i]
		}
	}
	return nil
}

// GetBankCardList Get bank card from storage.
func (r *Repo) GetBankCardList() []entity.BankCard {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.record.BankCardList
}

// GetBankCardByNumber Get bank card from storage by number.
func (r *Repo) GetBankCardByNumber(number string) *entity.BankCard {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i := range r.record.BankCardList {
		if r.record.BankCardList[i].Number == number {
			return &r.record.BankCardList[i]
		}
	}
	return nil
}
