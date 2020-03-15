package main

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

type rki struct {
	table     *goquery.Selection
	timestamp time.Time
	url       string
}

func cellFirstNumber(row *goquery.Selection, number int) int {
	return toNumber(strings.Split(row.Eq(number).Text(), " ")[0])
}

func loadRKI() (r rki) {
	c := colly.NewCollector()
	r.url = "https://www.rki.de/DE/Content/InfAZ/N/Neuartiges_Coronavirus/Fallzahlen.html"
	c.OnHTML("table", func(e *colly.HTMLElement) {
		r.table = e.DOM
	})
	c.OnHTML(".text > p.null", func(e *colly.HTMLElement) {
		// Stand: 14.3.2020, 15:00 Uhr (online aktualisiert um 20:00 Uhr)
		r.timestamp = toDate(e.DOM.Text())
	})
	c.Visit(r.url)
	return
}

func (r rki) lookup(region string) int {
	row := r.table.Find("tr").FilterFunction(func(i int, tr *goquery.Selection) bool {
		return tr.Find("td:first-child").Text() == region
	}).First().Find("td")
	conventional := cellFirstNumber(row, 1)
	electronic := cellFirstNumber(row, 2)
	if conventional > electronic {
		return conventional
	}
	return electronic
}
