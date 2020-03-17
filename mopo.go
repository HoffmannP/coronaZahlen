package main

import (
	"bytes"
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
)

type mopo struct {
	table     map[string]casecount
	timestamp time.Time
	url       string
}

func loadMopo() (m mopo) {
	c := colly.NewCollector()
	m.url = "https://interaktiv.morgenpost.de/corona-virus-karte-infektionen-deutschland-weltweit"
	m.table = make(map[string]casecount)
	date := int64(0)
	c.OnResponse(func(r *colly.Response) {
		lines := bytes.Split(r.Body, []byte("\n"))
		for _, line := range lines {
			if bytes.Equal(line[:12], []byte("Deutschland,")) {
				cells := bytes.Split(line, []byte(","))
				ts, _ := strconv.ParseInt(string(cells[2]), 10, 64)
				count, _ := strconv.Atoi(string(cells[4]))
				m.table[string(cells[1])] = casecount{
					URL:   string(cells[10]),
					Date:  ts / 1000,
					Count: count,
				}
				if ts > date {
					date = ts
				}
			}
		}
	})
	c.Visit(m.url + "/data/Coronavirus.current.v2.csv?" + strconv.Itoa(int(m.timestamp.Unix())))
	m.timestamp = time.Unix(date/1000, 0)
	return
}

func (m *mopo) lookup(region string) int {
	for regionName, regionData := range m.table {
		if regionName == region {
			return regionData.Count
		}
	}
	return -1
}
