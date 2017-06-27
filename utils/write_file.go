package utils

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func WriteFile(rootapi interface{}, basepath string) error {
	dt, err := json.MarshalIndent(rootapi, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(basepath, "swagger.json"), append(dt, '\n'), 0555)
	if err != nil {
		return err
	}
	dtyml, err := yaml.Marshal(rootapi)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(basepath, "swagger.yml"), dtyml, 0555)
	if err != nil {
		return err
	}
	return nil
}
