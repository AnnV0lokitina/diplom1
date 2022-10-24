package filestorage

import (
	"encoding/json"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	"io/ioutil"
)

// SetParamsFromJSON Sets configuration from json file.
func SetParamsFromJSON(path string, p *entity.Params) error {
	fContent, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(fContent, p); err != nil {
		return err
	}
	return nil
}
