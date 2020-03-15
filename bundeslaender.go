package main

import "github.com/gocolly/colly/v2"

type caseRegions map[string]caseRegion

type caseRegion struct {
	URL       string
	Listentry position
	Casecount position
	Timestamp position
}

type position struct {
	Selector string
	Match    string
}

func updateURL(regionData caseRegion) caseRegion {
	c := colly.NewCollector()
	c.OnHTML(regionData.Listentry.Selector, func(e *colly.HTMLElement) {
		path, exists := e.DOM.Attr("href")
		if exists {
			regionData.URL = regionData.Listentry.Match + path
		}
	})
	c.Visit(regionData.URL)
	return regionData
}

func regions() caseRegions {
	r := caseRegions{
		"Baden-Württemberg": {
			URL:       "https://sozialministerium.baden-wuerttemberg.de/de/gesundheit-pflege/gesundheitsschutz/infektionsschutz-hygiene/informationen-zu-coronavirus/",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: "figcaption",
				Match:    `([.\d]+) bestätigte Corona-Fälle`,
			},
			Timestamp: position{
				Selector: ".article__section > .text > .bodytext",
				Match:    "Stand: 2. Januar 2006",
			},
		},
		"Bayern": {
			URL:       "https://www.lgl.bayern.de/gesundheit/infektionsschutz/infektionskrankheiten_a_z/coronavirus/karte_coronavirus/index.htm",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: "div.row > div:nth-child(2) > table > tbody > tr:nth-child(9) > td:nth-child(2)",
				Match:    `([.\d]+)`,
			},
			Timestamp: position{
				Selector: ".container > .bildunterschrift",
				Match:    "*Stand: 2.01.2006 15:04 Uhr.",
			},
		},
		"Berlin": {
			URL:       "https://www.berlin.de/sen/gpg/service/presse/",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: ".rss > .html5-section > ul > li > a",
				Match:    `Coronavirus: Derzeit ([.\d]+) bestätigte Fälle`,
			},
			Timestamp: position{
				Selector: ".rss > .html5-section > ul > li > span",
				Match:    "2.01.2006",
			},
		},
		"Brandenburg": {
			URL:       "https://msgiv.brandenburg.de/msgiv/de/start/",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: "#form_id32516 > div > div > h2",
				Match:    `Insgesamt ([.\d]+) Erkrankungen`,
			},
			Timestamp: position{
				Selector: "#form_id32516 > div > div > p",
				Match:    " (Stand: 2.01.2006, 15:04 Uhr).",
			},
		},
		"Bremen": {
			URL:       "https://www.gesundheit.bremen.de/sixcms/detail.php?gsid=bremen229.c.32657.de",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: "#main > div:nth-child(3) > div",
				Match:    ` ([.\d]+) als positiv gemeldete`,
			},
			Timestamp: position{
				Selector: "#main > div:nth-child(3) > div",
				Match:    " (Stand 2.01.2006).",
			},
		},
		"Hamburg": {
			URL:       "https://www.hamburg.de/bgv/pressemeldungen/13719254/2020-03-13-bgv-coronavirus-aktuell/",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: "#pushedContainer > main > div.article > div:nth-child(7) > div > div > div",
				Match:    `insgesamt ([.\d]+) positiv Getestete`,
			},
			Timestamp: position{
				Selector: ".article-date",
				Match:    "2. Januar 2006 15:04 Uhr",
			},
		},
		"Hessen": {
			URL:       "https://soziales.hessen.de/gesundheit/infektionsschutz/coronavirus-sars-cov-2",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: "section:nth-child(2) > .block-inner > .blockContent",
				Match:    `insgesamt ([.\d]+) SARS-CoV-2-Fälle`,
			},
			Timestamp: position{
				Selector: "section:nth-child(2) > .block-inner > .blockContent",
				Match:    "Stand 2. Januar 2006,",
			},
		},
		"Mecklenburg-Vorpommern": {
			URL: "https://www.regierung-mv.de/Landesregierung/wm/Aktuell/?sa.query=neue+Corona-Infektionen&sa.pressemitteilungen.area=11",
			Listentry: position{
				Selector: ".resultlist > .teaser:nth-of-type(2) > div > h3 > a",
				Match:    "https://www.regierung-mv.de",
			},
			Casecount: position{
				Selector: ".dvz-contenttype-presseserviceassistent",
				Match:    `Insgesamt wurden bislang ([.\d]+) Menschen`,
			},
			Timestamp: position{
				Selector: ".dtstart",
				Match:    "2.01.2006",
			},
		},
		"Niedersachsen": {
			URL:       "https://www.niedersachsen.de/Coronavirus",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: ".maincontent > .content > .complementary > div",
				Match:    ` ([.\d]+)\s+laborbestätigte Covid-19-Fälle`,
			},
			Timestamp: position{
				Selector: ".maincontent > .content > .complementary > div > p > i",
				Match:    "zuletzt aktualisiert am 2.01.2006, 15 Uhr",
			},
		},
		"Nordrhein-Westfalen": {
			URL:       "https://www.mags.nrw/coronavirus-fallzahlen-nrw",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: "table > tbody > tr:nth-child(54) > td:nth-child(2)",
				Match:    `([.\d]+)`,
			},
			Timestamp: position{
				Selector: ".field-item > p",
				Match:    "Aktueller Stand der Liste: 2. Januar 20206, 15.04 Uhr.",
			},
		},
		"Rheinland-Pfalz": {
			URL:       "https://msagd.rlp.de/de/service/presse/detail/news/News/detail/information-der-landesregierung-zum-aktuellen-stand-hinsichtlich-des-coronavirus-bundesratsinitiati/",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: "h6",
				Match:    `insgesamt ([.\d]+) bestätigte SARS-CoV-2 Fälle`,
			},
			Timestamp: /* []position{ */
			position{
				Selector: ".news-list-date",
				Match:    "2.01.2006",
			},
			/*
					{
						Selector: ".news-text-wrap > p",
						Match:    "Stand: 15:04 Uhr",
					},
				},
			*/
		},
		"Saarland": {
			URL:       "https://www.saarland.de/SID-C29AF463-5CFEAE1B/253741.htm",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: ".textchapter_frame > p > strong",
				Match:    `landesweit auf ([.\d]+) – `,
			},
			Timestamp: position{
				Selector: ".textchapter_frame > p",
				Match:    "2.01.2006 (15:04)",
			},
		},
		"Sachsen": {
			URL:       "https://www.sms.sachsen.de/coronavirus.html",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: "table > tbody > tr:nth-child(14) > td:nth-child(2)",
				Match:    `([.\d]+)`,
			},
			Timestamp: position{
				Selector: ".text-col > p",
				Match:    "Stand: 2. Januar 2006, 15:04 Uhr",
			},
		},
		"Sachsen-Anhalt": {
			URL: "https://ms.sachsen-anhalt.de/presse/pressemitteilungen/?no_cache=1",
			Listentry: position{
				".tx-rssdisplay > div > div > h2 > a",
				"",
			},
			Casecount: position{
				Selector: ".pm_titel",
				Match:    `steigt auf ([.\d]+)`,
			},
			Timestamp: position{
				Selector: ".pm_datum",
				Match:    "Magdeburg, den 2. Januar 2006",
			},
		},
		"Schleswig-Holstein": {
			URL: "https://schleswig-holstein.de/SiteGlobals/Forms/Suche/DE/Expertensuche_Formular.html?gts=%2526aa42becc-5285-4475-af4d-0413dde5e634_list%253DdateOfIssue_dt%252Bdesc&documentType_=pressrelease&templateQueryString=Gesundheitsministerium+informiert+zum+Corona-Virus",
			Listentry: position{
				Selector: "ol#searchResult > li:first-child > a",
				Match:    "https://schleswig-holstein.de",
			},
			Casecount: position{
				Selector: "#content > .singleview > .bodyText",
				Match:    `positiv getesteten Covid19-Fälle in Schleswig-Holstein: Summe ([.\d]+)`,
			},
			Timestamp: position{
				Selector: ".docData > .value",
				Match:    "2.01.2006",
			},
		},
		"Thüringen": {
			URL:       "https://www.tmasgff.de/covid-19",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: "h3",
				Match:    `([.\d]+) bestätigte Infektionen`,
			},
			Timestamp: position{
				Selector: "h3",
				Match:    "(Stand 14.03.2006 - 15:04 Uhr)",
			},
		},
	}

	for regionName, regionData := range r {
		if regionData.Listentry.Selector != "" {
			r[regionName] = updateURL(regionData)
		}
	}

	return r
}
