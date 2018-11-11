package main

import (
	"encoding/json"
	"io/ioutil"
)

type student struct {
	Name string
	Sex  string
	Age  int
}

func (p *student) Save() (err error) {
	data, err := json.Marshal(p)
	if err != nil {
		return
	}

	err = ioutil.WriteFile("d:/hello.dat", data, 0755)
	return
}

func (p *student) Load() (err error) {
	data, err := ioutil.ReadFile("d:/hello.dat")
	if err != nil {
		return
	}
	json.Unmarshal(data, p)
	return
}
