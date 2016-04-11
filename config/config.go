package config

import (
	"encoding/json"	
	"io/ioutil"	
)

type Cmdconf struct {
	Username      string `json:"Username"`
	Password string `json:"Password"`
	Aliases   []*Alias `json:"storages"`
}

type Alias struct {
	Name          string `json:"Name"`
	Port          string `json:"Port"`
	Enabled       bool   `json:"enabled"`	
	Description   string `json:"description"`
}

func New(filename string) (result *Cmdconf, err error) {

	file, e := ioutil.ReadFile(filename)
	if e != nil {
		return nil, e
	}

	if e = json.Unmarshal(file, &result); e != nil {
		return nil, e
	}

	return result, nil
}