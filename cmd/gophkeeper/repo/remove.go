package repo

import "github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"

// RemoveTextFileByName Remove text file information from storage by name.
func (r *Repo) RemoveTextFileByName(name string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	var deletedItem *entity.File
	for i := range r.record.TextFileList {
		if r.record.TextFileList[i].Name == name {
			tmp := r.record.TextFileList[i]
			deletedItem = &tmp
			r.record.TextFileList = append(r.record.TextFileList[:i], r.record.TextFileList[i+1:]...)
			break
		}
	}
	if deletedItem == nil {
		return errorNotFound
	}
	err := r.enclosure.Remove(name)
	if err != nil {
		r.record.TextFileList = append(r.record.TextFileList, *deletedItem)
		return err
	}
	err = r.writer.WriteRecord(r.record)
	if err != nil {
		r.record.TextFileList = append(r.record.TextFileList, *deletedItem)
		return err
	}
	return nil
}

// RemoveBinaryFileByName Remove binary file information from storage by name.
func (r *Repo) RemoveBinaryFileByName(name string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	var deletedItem *entity.File
	for i := range r.record.BinaryFileList {
		if r.record.BinaryFileList[i].Name == name {
			tmp := r.record.BinaryFileList[i]
			deletedItem = &tmp
			r.record.BinaryFileList = append(r.record.BinaryFileList[:i], r.record.BinaryFileList[i+1:]...)
			break
		}
	}
	if deletedItem == nil {
		return errorNotFound
	}
	err := r.enclosure.Remove(name)
	if err != nil {
		r.record.BinaryFileList = append(r.record.BinaryFileList, *deletedItem)
		return err
	}
	err = r.writer.WriteRecord(r.record)
	if err != nil {
		r.record.BinaryFileList = append(r.record.BinaryFileList, *deletedItem)
		return err
	}
	return nil
}

// RemoveCredentialsByLogin Remove credentials information from storage by login.
func (r *Repo) RemoveCredentialsByLogin(login string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	var deletedItem *entity.Credentials
	for i := range r.record.CredentialsList {
		if r.record.CredentialsList[i].Login == login {
			tmp := r.record.CredentialsList[i]
			deletedItem = &tmp
			r.record.CredentialsList = append(r.record.CredentialsList[:i], r.record.CredentialsList[i+1:]...)
			break
		}
	}
	if deletedItem == nil {
		return errorNotFound
	}
	err := r.writer.WriteRecord(r.record)
	if err != nil {
		r.record.CredentialsList = append(r.record.CredentialsList, *deletedItem)
		return err
	}
	return nil
}

// RemoveBankCardByNumber Remove bank card information from storage by number.
func (r *Repo) RemoveBankCardByNumber(number string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	var deletedItem *entity.BankCard
	for i := range r.record.BankCardList {
		if r.record.BankCardList[i].Number == number {
			tmp := r.record.BankCardList[i]
			deletedItem = &tmp
			r.record.BankCardList = append(r.record.BankCardList[:i], r.record.BankCardList[i+1:]...)
			break
		}
	}
	if deletedItem == nil {
		return errorNotFound
	}
	err := r.writer.WriteRecord(r.record)
	if err != nil {
		r.record.BankCardList = append(r.record.BankCardList, *deletedItem)
		return err
	}
	return nil
}
