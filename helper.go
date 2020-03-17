package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/goodsign/monday"
)

func parseToRegexp(p string) string {
	p = regexp.MustCompile(`\b(?i)Monday\b`).ReplaceAllLiteralString(p, `%dayofweek%`)
	p = regexp.MustCompile(`\b0?1\b`).ReplaceAllLiteralString(p, `%month%`)
	p = regexp.MustCompile(`\b(?i)January\b`).ReplaceAllLiteralString(p, `%monthName%`)
	p = regexp.MustCompile(`\b0?2\b\b`).ReplaceAllLiteralString(p, `%day%`)
	p = regexp.MustCompile(`\b15|0?3\b`).ReplaceAllLiteralString(p, `%hour%`)
	p = regexp.MustCompile(`\b0?4\b`).ReplaceAllLiteralString(p, `%minute%`)
	p = regexp.MustCompile(`\b(20?)06\b`).ReplaceAllLiteralString(p, `%year%`)
	p = regexp.MustCompile(`%dayofweek%`).ReplaceAllLiteralString(p, `\b(?:Mo|Di|Mi|Do|Fr|Sa|So)[[:alpha:]]{0,8}\b`)
	p = regexp.MustCompile(`%month%`).ReplaceAllLiteralString(p, `\b(?:0?\d|1[012])\b`)
	p = regexp.MustCompile(`%monthName%`).ReplaceAllLiteralString(p, `\b[[:alpha:]ä]{3,9}\b`)
	p = regexp.MustCompile(`%day%`).ReplaceAllLiteralString(p, `\b(?:[012]?\d|3[01])\b`)
	p = regexp.MustCompile(`%hour%`).ReplaceAllLiteralString(p, `\b(?:[01]?\d|2[0-4])\b`)
	p = regexp.MustCompile(`%minute%`).ReplaceAllLiteralString(p, `\b(?:[0-5]?\d|60)\b`)
	return "(" + regexp.MustCompile(`%year%`).ReplaceAllLiteralString(p, `\b(?:\d\d)?\d\d\b`) + ")"
}

func errorHandler(text string, err error) {
	fmt.Printf("Error parsing '%s'\n", text)
	panic(err)
}

func (p position) grabNumber(e *colly.HTMLElement) int {
	return toNumber(grab(e, p))
}

func toNumber(t string) int {
	if t == "" {
		return -1
	}
	text := strings.ReplaceAll(t, ".", "")
	i, err := strconv.Atoi(text)
	if err != nil {
		errorHandler(text, err)
	}
	return int(i)
}

func (p position) grabDate(e *colly.HTMLElement) time.Time {
	layout := p.Match
	p.Match = parseToRegexp(regexp.QuoteMeta(p.Match))
	return toDate(grab(e, p), layout)
}

func toDate(ts, layout string) time.Time {
	if ts == "" {
		return time.Unix(0, 0)
	}
	tz, _ := time.LoadLocation("Europe/Berlin")
	t, err := monday.ParseInLocation(layout, ts, tz, "de_DE")
	if err != nil {
		errorHandler(ts, err)
	}
	if t.Year() == 0 {
		t = t.AddDate(time.Now().Year(), 0, 0)
	}
	return t
}

func grab(e *colly.HTMLElement, p position) string {
	re := regexp.MustCompile(p.Match)
	var t []string
	matches := e.DOM.Find(p.Selector)
	if matches.Length() == 0 {
		panic("Konnte kein Element mit '" + p.Selector + "' finden")
	}
	matches.EachWithBreak(func(i int, s *goquery.Selection) bool {
		t = re.FindStringSubmatch(s.Text())
		return len(t) == 0 // i. e. while no submatch found continue
	})
	if len(t) <= 1 {
		panic("Konnte nichts matchen mit '" + p.Match + "'")
	}
	return t[1]
}

func (r *caseRegion) url() string {
	if r.Listentry.Selector != "" {
		c := colly.NewCollector()
		c.OnHTML(r.Listentry.Selector, func(e *colly.HTMLElement) {
			path, exists := e.DOM.Attr("href")
			if exists {
				r.URL = r.Listentry.Match + path
			}
		})
		c.Visit(r.URL)
	}
	return r.URL
}

func (r *caseRegion) loadRegion() (num int, ts time.Time) {
	if r.Casecount.Selector == "" {
		return -1, time.Unix(0, 0)
	}
	c := colly.NewCollector()
	c.SetRequestTimeout(15 * 1000000000)
	c.OnHTML("body", func(e *colly.HTMLElement) {

		num = r.Casecount.grabNumber(e)
		ts = r.Timestamp.grabDate(e)
	})
	c.OnError(func(r *colly.Response, err error) {
		panic(err)
	})
	c.Visit(r.url())
	return
}
