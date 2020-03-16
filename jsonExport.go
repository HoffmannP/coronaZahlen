package main

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type casecount struct {
	URL       string
	Timestamp int64
	Count     int
	RKI       int
}

type jsonData struct {
	name   string
	counts map[string]casecount
}

func newJSON(name string) (j jsonData) {
	j.name = name
	j.counts = make(map[string]casecount)
	return
}

func (j *jsonData) getName() string {
	return j.name
}

func (j *jsonData) append(name, url string, timestamp time.Time, count, rki int) {
	j.counts[name] = casecount{
		URL:       url,
		Timestamp: timestamp.Unix(),
		Count:     count,
		RKI:       rki,
	}
}

func (j *jsonData) save(path string) {
	json, err := json.MarshalIndent(j.counts, "", "  ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(path+j.name+".json", json, 0644)
	if err != nil {
		panic(err)
	}
}
