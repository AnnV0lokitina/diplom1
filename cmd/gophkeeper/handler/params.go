package handler

import (
	"encoding/json"
	"io/ioutil"
)

type Params struct {
	ServerAddress string `json:"server_address,omitempty"`
	Login         string `json:"login,omitempty"`
	Password      string `json:"password,omitempty"`
	FileStorePath string `json:"file_store_path,omitempty"`
}

func (p *Params) SetFromJSON(path string) error {
	fContent, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(fContent, p); err != nil {
		return err
	}
	return nil
}
