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

func getFtpData(filename string) (ftpdata ftpdata) {
	j, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(j, &ftpdata); err != nil {
		panic(err)
	}
	return
}

func connect(ftpdata ftpdata) func(*bytes.Buffer) error {
	c, err := ftp.Dial(ftpdata.Address)
	if err != nil {
		panic(err)
	}

	c.Login(ftpdata.User, ftpdata.Password)
	if err != nil {
		panic(err)
	}

	return func(data *bytes.Buffer) error {
		if err := c.Stor(ftpdata.Remotename, data); err != nil {
			return err
		}

		if err := c.Quit(); err != nil {
			return err
		}

		return nil
	}
}

func upload(file []byte) {
	data := bytes.NewBuffer(file)
	upload := connect(getFtpData("ftp.json"))
	if err := upload(data); err != nil {
		panic(err)
	}
}
