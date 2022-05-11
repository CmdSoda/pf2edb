package pf2e

import (
	"encoding/json"
	"io/ioutil"
)

type Configuration struct {
	DataFolder        string
	SystemName        string
	TranslationModule string
}

func NewConfiguration() Configuration {
	return Configuration{
		DataFolder:        "",
		SystemName:        "",
		TranslationModule: "",
	}
}

func (c *Configuration) Load(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, c)
	if err != nil {
		return err
	}
	return nil
}

func (c Configuration) Save(filename string) error {
	buffer, err := json.Marshal(c)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, buffer, 0644)
	if err != nil {
		return err
	}
	return nil
}
