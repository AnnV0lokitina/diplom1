package entity

// Credentials with Login - Password pair.
type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Meta     string `json:"meta"`
}

// BankCard with bank card data.
type BankCard struct {
	Number     string `json:"number"`
	ExpDate    string `json:"exp_date"`
	Cardholder string `json:"cardholder"`
	Code       string `json:"code"`
	Meta       string `json:"meta"`
}

type File struct {
	Name string `json:"name"`
	Meta string `json:"meta"`
}

// Record with user Information.
type Record struct {
	CredentialsList []Credentials `json:"login_password_list"`
	TextFileList    []File        `json:"text_file_list"`
	BinaryFileList  []File        `json:"binary_file_list"`
	BankCardList    []BankCard    `json:"bank_card_list"`
}
