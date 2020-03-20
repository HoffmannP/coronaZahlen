package main

import (
	"time"

	"github.com/gocolly/colly/v2"
)

type rkiType struct {
	counts    map[string]int
	timestamp time.Time
	url       string
}

func loadRKI() (rki rkiType, err error) {
	rki.url = "https://www.rki.de/DE/Content/InfAZ/N/Neuartiges_Coronavirus/Fallzahlen.html"
	c := colly.NewCollector()
	c.OnHTML("#main > .text", func(e *colly.HTMLElement) {
		rki.counts, err = rki.count(e)
		if err != nil {
			return
		}
		rki.timestamp, err = rki.date(e)
		return
	})
	c.Visit(rki.url)
	return
}

func (rki *rkiType) count(e *colly.HTMLElement) (counts map[string]int, err error) {
	counts = make(map[string]int)
	rows := e.DOM.Find("table > tbody > tr")
	for i := 0; i < rows.Length(); i++ {
		cells := rows.Eq(i).Find("td")
		counts[cells.Eq(0).Text()], err = toNumber(cells.Eq(1).Text())
	}
	return
}

func (rki *rkiType) date(e *colly.HTMLElement) (time.Time, error) {
	return position{
		Selector: "p.null",
		Match:    "Stand: 2.1.2006, 15:04 Uhr",
	}.grabDate(e)
}

func (rki *rkiType) lookup(region string) int {
	return rki.counts[region]
}
