package models

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func ReadYAMLFile(fileName string, object interface{}) error {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(content, object)
}
