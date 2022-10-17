package handler

import (
	"encoding/json"
	"io/ioutil"
)

// Params Contains information about application launch options.
type Params struct {
	ServerAddress string `json:"server_address,omitempty"`
	FileStorePath string `json:"file_store_path,omitempty"`
	ArchiveName   string `json:"archive_name,omitempty"`
}

// SetFromJSON Sets configuration from json file.
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
