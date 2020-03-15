package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

func errorHandler(text string, err error) {
	fmt.Printf("Error parsing '%s'\n", text)
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

func toDate(ts string) time.Time {
	if ts == "" {
		return time.Unix(0, 0)
	}
	t, err := time.Parse("01.02.2006", ts)
	if err != nil {
		errorHandler(ts, err)
	}
	return t
}

func grab(e *colly.HTMLElement, p position) string {
	re := regexp.MustCompile(p.Match)
	var t []string
	e.DOM.Find(p.Selector).EachWithBreak(func(i int, s *goquery.Selection) bool {
		t = re.FindStringSubmatch(s.Text())
		return len(t) == 0 // i. e. no submatch found
	})
	if len(t) == 0 {
		panic("Konnte nichts matchen mit " + p.Match)
	}
	return t[1]
}

func loadRegion(r caseRegion) (num int, ts time.Time) {
	if r.Casecount.Selector == "" {
		return -1, time.Unix(0, 0)
	}
	c := colly.NewCollector()
	c.OnHTML("body", func(e *colly.HTMLElement) {
		num = toNumber(grab(e, r.Casecount))
		ts = toDate(grab(e, r.Timestamp))
	})
	c.Visit(r.URL)
	return
}
