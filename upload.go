package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/jlaffaye/ftp"
)

type ftpdata struct {
	Address    string
	User       string
	Password   string
	Remotename string
}

func upload(file []byte) {
	j, err := ioutil.ReadFile("ftp.json")
	if err != nil {
		panic(err)
	}

	var ftpdata ftpdata
	if err := json.Unmarshal(j, &ftpdata); err != nil {
		panic(err)
	}

	c, err := ftp.Dial(ftpdata.Address)
	if err != nil {
		panic(err)
	}

	c.Login(ftpdata.User, ftpdata.Password)
	if err != nil {
		panic(err)
	}

	data := bytes.NewBuffer(file)
	err = c.Stor(ftpdata.Remotename, data)
	if err != nil {
		panic(err)
	}

	if err := c.Quit(); err != nil {
		panic(err)
	}
}
