package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
	Mopo  int
	CJ    int
}

type casecounts map[string]casecountP

type data struct {
	Name    string
	Date    int64
	RKI     casecount
	Mopo    casecount
	CJ      casecount
	Regions casecounts
	Sum     int
}

func newData(name string, rki rkiType, m mopo, cj cjType) (j data) {
	j.Name = name
	j.RKI = casecount{
		URL:   rki.url,
		Date:  rki.timestamp.Unix(),
		Count: rki.lookup("Gesamt"),
	}
	j.Mopo = casecount{
		URL:   m.url,
		Date:  m.timestamp.Unix(),
		Count: -1,
	}
	j.CJ = casecount{
		URL:   cj.url,
		Date:  cj.timestamp.Unix(),
		Count: -1,
	}
	j.Regions = make(casecounts)
	return
}

func (j *data) append(name, url string, timestamp time.Time, count, rki, mopo, cj int) {
	j.Regions[name] = casecountP{
		URL:   url,
		Date:  timestamp.Unix(),
		Count: count,
		RKI:   rki,
		Mopo:  mopo,
		CJ:    cj,
	}
}

func (j *data) sum(sum int) {
	j.Sum = sum
	j.Date = time.Now().Unix()
}

func (j *data) saveJSON(filename string) {
	if !diff(filename, &j.Regions) {
		return
	}
	log.Println("Update")

	json, err := json.MarshalIndent(j, "", "  ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, json, 0644)
	if err != nil {
		panic(err)
	}
	upload(json)
}
