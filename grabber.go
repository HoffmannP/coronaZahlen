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
	p = regexp.MustCompile(`.\b0?4\b`).ReplaceAllLiteralString(p, `%minute%`)
	p = regexp.MustCompile(`\b(20?)06\b`).ReplaceAllLiteralString(p, `%year%`)
	p = regexp.QuoteMeta(p)
	p = regexp.MustCompile(`%dayofweek%`).ReplaceAllLiteralString(p, `\b(?:Mo|Di|Mi|Do|Fr|Sa|So)[[:alpha:]]{0,8}\b`)
	p = regexp.MustCompile(`%month%`).ReplaceAllLiteralString(p, `\b(?:0?\d|1[012])\b`)
	p = regexp.MustCompile(`%monthName%`).ReplaceAllLiteralString(p, `\b[[:alpha:]ä]{3,9}\b`)
	p = regexp.MustCompile(`%day%`).ReplaceAllLiteralString(p, `\b(?:[012]?\d|3[01])\b`)
	p = regexp.MustCompile(`%hour%`).ReplaceAllLiteralString(p, `\b(?:[01]?\d|2[0-4])\b`)
	p = regexp.MustCompile(`%minute%`).ReplaceAllLiteralString(p, `(?:.%minute%)?`)
	p = regexp.MustCompile(`%minute%`).ReplaceAllLiteralString(p, `\b(?:[0-5]?\d|60)\b`)
	return "(" + regexp.MustCompile(`%year%`).ReplaceAllLiteralString(p, `\b(?:\d\d)?\d\d\b`) + ")"
}

func errorHandler(text string, err error) {
	fmt.Printf("Error parsing '%s'\n", text)
	panic(err)
}

func (p position) grabNumber(e *colly.HTMLElement) (int, error) {
	numtext, err := grab(e, p)
	if err != nil {
		return -1, err
	}
	return toNumber(numtext)
}

func toNumber(t string) (int, error) {
	if t == "" {
		return -1, fmt.Errorf("Keine Zahltext zum Umwandeln")
	}
	text := strings.ReplaceAll(t, ".", "")
	i, err := strconv.Atoi(text)
	if err != nil {
		return -1, err
	}
	return int(i), nil
}

func (p position) grabDate(e *colly.HTMLElement) (time.Time, error) {
	layout := p.Match
	p.Match = parseToRegexp(p.Match)
	date, err := grab(e, p)
	if err != nil {
		return time.Time{}, err
	}
	return toDate(date, layout)
}

func toDate(ts, layout string) (time.Time, error) {
	if ts == "" {
		return time.Time{}, fmt.Errorf("no count selector")
	}
	ts = strings.Replace(ts, "onnabend", "amstag", 1)
	ts = strings.Replace(ts, "iensta,", "ienstag,", 1)
	tz, _ := time.LoadLocation("Europe/Berlin")
	t, err := monday.ParseInLocation(layout, ts, tz, "de_DE")
	if err != nil {
		layout = regexp.MustCompile(`.04`).ReplaceAllLiteralString(layout, "")
		t, err = monday.ParseInLocation(layout, ts, tz, "de_DE")
	}
	if err != nil {
		layout = regexp.MustCompile(`:04`).ReplaceAllLiteralString(layout, "")
		t, err = monday.ParseInLocation(layout, ts, tz, "de_DE")
	}
	if err != nil {
		return time.Time{}, err
	}
	if t.Year() == 0 {
		t = t.AddDate(time.Now().Year(), 0, 0)
	}
	return t, nil
}

func grab(e *colly.HTMLElement, p position) (string, error) {
	re := regexp.MustCompile(p.Match)
	var t []string
	matches := e.DOM.Find(p.Selector)
	if matches.Length() == 0 {
		return "", fmt.Errorf("Konnte kein Element mit '" + p.Selector + "' finden")
	}
	matches.EachWithBreak(func(i int, s *goquery.Selection) bool {
		t = re.FindStringSubmatch(s.Text())
		return len(t) == 0 // i. e. while no submatch found continue
	})
	if len(t) <= 1 {
		return "", fmt.Errorf("Konnte mit '%s' nicht matchen", p.Match)
	}
	return t[1], nil
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

func (r *caseRegion) loadRegion() (num int, ts time.Time, err error) {
	if r.Casecount.Selector == "" {
		return -1, time.Time{}, nil
	}
	maxTries := 2
	retries := 0
	c := colly.NewCollector()
	c.OnHTML("body", func(e *colly.HTMLElement) {
		num, err = r.Casecount.grabNumber(e)
		if err != nil {
			return
		}
		ts, err = r.Timestamp.grabDate(e)
		if err != nil {
			return
		}
	})
	c.OnError(func(p *colly.Response, netErr error) {
		if retries < maxTries {
			retries = retries + 1
			err = fmt.Errorf("Needed %d retries b/c timeout", retries)
			c.Visit(r.url())
		} else {
			err = netErr
		}
	})
	c.Visit(r.url())
	return
}
