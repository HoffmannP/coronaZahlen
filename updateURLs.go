package main

import (
	"github.com/gocolly/colly/v2"
)

func updateURLs(r caseRegions) caseRegions {
	c := colly.NewCollector()

	MV := r["Mecklenburg-Vorpommern"]
	c.OnHTML(".resultlist > .teaser:nth-of-type(2) > div > h3 > a", func(e *colly.HTMLElement) {
		path, exists := e.DOM.Attr("href")
		if exists {
			MV.URL = "https://www.regierung-mv.de" + path
		}
	})
	c.Visit(MV.URL)
	r["Mecklenburg-Vorpommern"] = MV

	SH := r["Schleswig-Holstein"]
	c.OnHTML("ol#searchResult > li:first-child > a", func(e *colly.HTMLElement) {
		path, exists := e.DOM.Attr("href")
		if exists {
			SH.URL = "https://schleswig-holstein.de/" + path
		}
	})
	c.Visit(SH.URL)
	r["Schleswig-Holstein"] = SH
	return r
}
