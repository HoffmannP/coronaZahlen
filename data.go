package main

import (
	"encoding/json"
	"io/ioutil"
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
	Max   int
}

type casecounts map[string]casecountP

type data struct {
	Name    string
	Date    int64
	RKI     casecount
	Mopo    casecount
	Regions casecounts
	Sum     int
}

func newData(name string, rki rkiType, m mopo) (j data) {
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
	j.Regions = make(casecounts)
	return
}

func (j *data) append(name, url string, timestamp time.Time, count, rki, mopo int) {
	var max int
	if count >= rki && count >= mopo {
		max = count
	} else if rki >= count && rki >= mopo {
		max = rki
	} else {
		max = mopo
	}
	j.Regions[name] = casecountP{
		URL:   url,
		Date:  timestamp.Unix(),
		Count: count,
		RKI:   rki,
		Mopo:  mopo,
		Max:   max,
	}
}

func (j *data) sum(sum int) {
	j.Sum = sum
	j.Date = time.Now().Unix()
}

func (j *data) saveJSON(filename string) {
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
