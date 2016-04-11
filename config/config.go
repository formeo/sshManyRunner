package config

import (
	"encoding/json"	
	"io/ioutil"	
)

type MyJsonName struct {
	Cmdconf struct {
		Aliases []struct {
			Description string `json:"Description"`
			Enabled     bool   `json:"Enabled"`
			Name        string `json:"Name"`
			Port        string `json:"Port"`
		} `json:"Aliases"`
		Password string `json:"Password"`
		Username string `json:"Username"`
	} `json:"Cmdconf"`
}


func New(filename string) (result Cmdconf, err error) {
   
	file, e := ioutil.ReadFile(filename)
	if e != nil {
		return result, e
	}

	if e = json.Unmarshal(file, &result); e != nil {
		return result, e
	}

	return result, nil
}