package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"text/template"
	"time"
)

type casecount struct {
	URL   string
	Date  int64
	Count int
}

type casecountP struct {
	URL   string
	Date  int64
	Count int
	RKI   int
}

type casecounts map[string]casecountP

type data struct {
	Name    string
	Date    int64
	RKI     casecount
	Regions casecounts
	Sum     int
}

func newData(name string, rki rkiType) (j data) {
	j.Name = name
	j.RKI = casecount{
		URL:   rki.url,
		Date:  rki.timestamp.Unix(),
		Count: rki.lookup("Gesamt"),
	}
	j.Regions = make(casecounts)
	return
}

func (j *data) append(name, url string, timestamp time.Time, count, rki int) {
	j.Regions[name] = casecountP{
		URL:   url,
		Date:  timestamp.Unix(),
		Count: count,
		RKI:   rki,
	}
}

func (j *data) sum(sum int) {
	j.Sum = sum
	j.Date = time.Now().Unix()
}

func (j *data) saveJSON(path string) {
	json, err := json.MarshalIndent(j, "", "  ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(path+j.Name+".json", json, 0644)
	if err != nil {
		panic(err)
	}
}

func (j *data) saveHTML(tmplFile, filename string) {
	funcMap := template.FuncMap{
		"blass": func(lt bool) string {
			if lt {
				return "class=\"blass\""
			}
			return ""
		},
		"isAvailable": func(c int) string {
			if c == -1 {
				return "n/a"
			}
			return strconv.Itoa(c)
		},
		"timestamp": func(t int64) string {
			return time.Unix(t, 0).Format("2.01.2006 15:04 Uhr")
		},
	}

	t, err := template.New(tmplFile).Funcs(funcMap).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	t = t.Lookup(tmplFile)

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, j)
	if err != nil {
		panic(err)
	}
}
