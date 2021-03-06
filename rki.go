package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type rkiType struct {
	counts    map[string]int
	timestamp time.Time
	url       string
}

func loadRKI() (rki rkiType, err error) {
	found := false
	selector := "#main > .text"
	rki.url = "https://www.rki.de/DE/Content/InfAZ/N/Neuartiges_Coronavirus/Fallzahlen.html"
	c := colly.NewCollector()
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		found = true
		rki.counts, err = rki.count(e)
		if err != nil {
			return
		}
		rki.timestamp, err = rki.date(e)
		return
	})
	c.Visit(rki.url)
	if !found {
		err = fmt.Errorf("Selektor '%s' wurde nicht gefunden", selector)
	}
	return
}

func (rki *rkiType) count(e *colly.HTMLElement) (counts map[string]int, err error) {
	counts = make(map[string]int)
	rows := e.DOM.Find("table > tbody > tr")
	for i := 0; i < rows.Length(); i++ {
		cells := rows.Eq(i).Find("td")
		counts[strings.Replace(cells.Eq(0).Text(), "­", "", -1)], err = toNumber(cells.Eq(1).Text())
	}
	return
}

func (rki *rkiType) date(e *colly.HTMLElement) (time.Time, error) {
	return position{
		Selector: "p",
		Match:    "Stand: 2.1.2006, 15:04 Uhr",
	}.grabDate(e)
}

func (rki *rkiType) lookup(region string) int {
	return rki.counts[region]
}
