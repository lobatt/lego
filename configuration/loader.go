package configuration

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"errors"
)

func LoadFromJson(rawConfig []byte, conf interface{}) error {
	return json.Unmarshal(rawConfig, conf)
}


func LoadFromJsonFile(path string, conf interface{}) error {
	f, er := os.OpenFile(path, os.O_RDONLY, 0666)
	if er != nil {
		return errors.New("Could not open config file: " + er.Error())
	}

	b, e := ioutil.ReadAll(f)
	if e != nil {
		return e
	}
	return LoadFromJson(b, conf)
}


