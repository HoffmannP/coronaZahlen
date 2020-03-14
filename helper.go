package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

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
	text := strings.ReplaceAll(t, ".", "")
	i, err := strconv.ParseInt(text, 10, 0)
	if err != nil {
		errorHandler(text, err)
	}
	return int(i)
}

func loadFromText(url, path, match string) (text string) {
	if path == "" {
		return "0"
	}
	c := colly.NewCollector()
	c.OnHTML("body", func(e *colly.HTMLElement) {
		re := regexp.MustCompile(match)
		t := re.FindStringSubmatch(e.DOM.Find(path).First().Text())
		text = t[1]
	})
	c.Visit(url)
	return
}

func loadRegion(r caseRegion) int {
	return toNumber(loadFromText(r.URL, r.Selector, r.Match))
}
