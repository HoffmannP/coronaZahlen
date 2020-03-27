package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
)

type cjType struct {
	counts    map[string]int
	timestamp time.Time
	url       string
}

func loadCj() (cj cjType, err error) {
	cj.url = "https://www.coronavirus.jetzt/karten/deutschland/"
	cj.counts, err = cj.count()
	if err != nil {
		return
	}
	cj.timestamp, err = cj.date()
	return
}

func (cj *cjType) count() (counts map[string]int, err error) {
	counts = make(map[string]int)
	found := false
	selector := "table"
	c := colly.NewCollector()
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		found = true
		rows := e.DOM.Find("tr")
		for i := 0; i < rows.Length(); i++ {
			cells := rows.Eq(i).Find("td")
			counts[cells.Eq(0).Text()], err = toNumber(cells.Eq(1).Text())
		}
	})
	c.Visit(cj.url)
	if !found {
		err = fmt.Errorf("Selektor '%s' wurde nicht gefunden", selector)
	}
	return
}

func (cj *cjType) date() (date time.Time, err error) {
	c := colly.NewCollector()
	found := false
	selector := ".vc_row > .vc_col-sm-4:first-child > .vc_column-inner > div > div > p"
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		found = true
		date, err = toDate(e.DOM.Text(), "Letzte Aktualisierung aller Zahlen. 2.01.2006, 15.04 Uhr.")
	})
	c.Visit("https://www.coronavirus.jetzt/")
	if !found {
		err = fmt.Errorf("Selektor '%s' wurde nicht gefunden", selector)
	}
	return
}

func (cj *cjType) lookup(region string) int {
	return cj.counts[region]
}
