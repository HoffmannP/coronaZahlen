package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

type caseRegions map[string]caseRegion

type caseRegion struct {
	URL      string
	Selector string
	Match    string
}

func errorHandler(text string, err error) {
	fmt.Printf("Error parsing text '%s'\n", text)
	panic(err)
}

func toNumber(t string) int {
	if t == "" {
		return 0
	}
	text := strings.ReplaceAll(t, ".", "")
	i, err := strconv.ParseInt(text, 10, 0)
	if err != nil {
		errorHandler(text, err)
	}
	return int(i)
}

func loadRegion(r caseRegion) (num int) {
	if r.Selector == "" {
		return -1
	}
	c := colly.NewCollector()
	c.OnHTML("body", func(e *colly.HTMLElement) {
		re := regexp.MustCompile(r.Match)
		var t []string
		e.DOM.Find(r.Selector).EachWithBreak(func(i int, s *goquery.Selection) bool {
			t = re.FindStringSubmatch(s.Text())
			return len(t) == 0 // i. e. no submatch found
		})
		if len(t) == 0 {
			panic("Konnte nichts matchen in " + r.URL)
		}
		num = toNumber(t[1])
	})
	c.Visit(r.URL)
	return
}
