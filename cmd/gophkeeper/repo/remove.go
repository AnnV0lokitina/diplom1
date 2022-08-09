package repo

// RemoveTextFileByName Remove text file information from storage by name.
func (r *Repo) RemoveTextFileByName(name string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i := range r.record.TextFileList {
		if r.record.TextFileList[i].Name == name {
			r.record.TextFileList = append(r.record.TextFileList[:i], r.record.TextFileList[i+1:]...)
			break
		}
	}
	return r.writer.WriteRecord(r.record)
}

// RemoveBinaryFileByName Remove binary file information from storage by name.
func (r *Repo) RemoveBinaryFileByName(name string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i := range r.record.BinaryFileList {
		if r.record.BinaryFileList[i].Name == name {
			r.record.BinaryFileList = append(r.record.BinaryFileList[:i], r.record.BinaryFileList[i+1:]...)
			break
		}
	}
	return r.writer.WriteRecord(r.record)
}

// RemoveCredentialsByLogin Remove credentials information from storage by login.
func (r *Repo) RemoveCredentialsByLogin(login string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i := range r.record.CredentialsList {
		if r.record.CredentialsList[i].Login == login {
			r.record.CredentialsList = append(r.record.CredentialsList[:i], r.record.CredentialsList[i+1:]...)
			break
		}
	}
	return r.writer.WriteRecord(r.record)
}

// RemoveBankCardByNumber Remove bank card information from storage by number.
func (r *Repo) RemoveBankCardByNumber(number string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i := range r.record.BankCardList {
		if r.record.BankCardList[i].Number == number {
			r.record.BankCardList = append(r.record.BankCardList[:i], r.record.BankCardList[i+1:]...)
			break
		}
	}
	return r.writer.WriteRecord(r.record)
}
