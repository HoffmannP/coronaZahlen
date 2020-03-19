package main

import (
	"bytes"
	"os"

	"github.com/jlaffaye/ftp"
)

func upload(file []byte) {
	c, err := ftp.Dial("ichplatz.de")
	if err != nil {
		panic(err)
	}

	c.Login(os.Getenv("ftpuser"), os.Getenv("ftppass"))
	if err != nil {
		panic(err)
	}

	data := bytes.NewBuffer(file)
	err = c.Stor(os.Getenv("ftpfile"), data)
	if err != nil {
		panic(err)
	}

	if err := c.Quit(); err != nil {
		panic(err)
	}
}
