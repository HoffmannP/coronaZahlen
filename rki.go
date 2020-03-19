package main

import (
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

type rkiType struct {
	table     *goquery.Selection
	timestamp time.Time
	url       string
}

func loadRKI() (rki rkiType) {
	c := colly.NewCollector()
	rki.url = "https://www.rki.de/DE/Content/InfAZ/N/Neuartiges_Coronavirus/Fallzahlen.html"
	c.OnHTML("#main > .text", func(e *colly.HTMLElement) {
		rki.table = e.DOM.Find("table")
		ts, err := position{
			Selector: "p.null",
			Match:    "Stand: 2.1.2006, 15:04 Uhr",
		}.grabDate(e)
		if err != nil {
			log.Println(err.Error())
		} else {
			rki.timestamp = ts
		}
	})
	c.Visit(rki.url)
	return
}

func (rki rkiType) lookup(region string) int {
	row := rki.table.Find("tr").FilterFunction(func(i int, tr *goquery.Selection) bool {
		return tr.Find("td:first-child").Text() == region
	}).First()
	count, err := toNumber(strings.Split(row.Find("td").Eq(1).Text(), " ")[0])
	if err != nil {
		log.Println(err.Error())
		return -1
	}
	return count
}
