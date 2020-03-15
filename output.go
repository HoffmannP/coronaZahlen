package main

import (
	"os"
	"text/template"
	"time"
)

type table struct {
	Header []string
	Rows   [][]string
	Footer []string
	Date   string
}

func display(header []string, rows [][]string, footer []string, file string) {
	t, err := template.ParseFiles("template.html")
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	err = t.Execute(f, table{
		Header: header,
		Rows:   rows,
		Footer: footer,
		Date:   time.Now().Format("2.01.2006 15:04"),
	})
	if err != nil {
		panic(err)
	}
}
