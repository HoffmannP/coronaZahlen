package main

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

func regions() caseRegions {
	return caseRegions{
		"Baden-Württemberg": {
			URL:       "https://sozialministerium.baden-wuerttemberg.de/de/gesundheit-pflege/gesundheitsschutz/infektionsschutz-hygiene/informationen-zu-coronavirus/",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: "figcaption",
				Match:    `([.\d]+) bestätigte Corona-Fälle`,
			},
			Timestamp: position{
				Selector: ".article__section > .text > .bodytext",
				Match:    "Stand: 2. January 2006",
			},
		},
		"Bayern": {
			URL:       "https://www.lgl.bayern.de/gesundheit/infektionsschutz/infektionskrankheiten_a_z/coronavirus/karte_coronavirus/index.htm",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: "div.row > div:nth-child(2) > table > tbody > tr:last-child > td:last-child",
				Match:    `([.\d]+)`,
			},
			Timestamp: position{
				Selector: ".row > div > .bildunterschrift",
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
			URL:       "https://www.gesundheit.bremen.de/",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: ".news_bigteaser > p",
				Match:    ` ([.\d]+) als positiv`,
			},
			Timestamp: position{
				Selector: ".news_bigteaser > p",
				Match:    "(Stand 2.1.2006)",
			},
		},
		"Hamburg": {
			URL:       "https://www.hamburg.de/coronavirus/",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: ".c_chart.one > .c_chart_h2",
				Match:    `([.\d]+)`,
			},
			Timestamp: position{
				Selector: ".chart_publication",
				Match:    "Stand: Monday, 2. January 2006",
			},
		},
		"Hessen": {
			URL:       "https://soziales.hessen.de/gesundheit/infektionsschutz/coronavirus-sars-cov-2",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: "section:nth-child(2) > .block-inner > .blockContent",
				Match:    `insgesamt ([.\d]+).SARS-CoV-2-Fälle`,
			},
			Timestamp: position{
				Selector: "section:nth-child(2) > .block-inner > .blockContent",
				Match:    "Stand 2. January 2006,",
			},
		},
		"Mecklenburg-Vorpommern": {
			URL: "https://www.regierung-mv.de/Landesregierung/wm/Aktuell/?sa.query=neue+Corona-Infektionen&sa.pressemitteilungen.area=11",
			Listentry: position{
				Selector: ".resultlist > .teaser:nth-of-type(2) > div > h3 > a",
				Match:    "https://www.regierung-mv.de",
			},
			Casecount: position{
				Selector: "", // table > tbody > tr:last-child > td:nth-child(3)",
				Match:    `([.\d]+)`,
			},
			Timestamp: position{
				Selector: "table > tbody > tr:first-child > td:nth-child(3)",
				Match:    "2.01. 15:04",
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
				Selector: ".maincontent > .content > .complementary > div",
				Match:    "aktualisiert am 2.01.2006, 15.04 Uhr",
			},
		},
		"Nordrhein-Westfalen": {
			URL:       "https://www.mags.nrw/coronavirus-fallzahlen-nrw",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: "table > tbody > tr:last-child > td:last-child",
				Match:    `([.\d]+)`,
			},
			Timestamp: position{
				Selector: ".field-item > p",
				Match:    "Aktueller Stand der Liste: 2. January 2006, 15.04 Uhr.",
			},
		},
		"Rheinland-Pfalz": {
			URL:       "https://msagd.rlp.de/de/unsere-themen/gesundheit-und-pflege/gesundheitliche-versorgung/oeffentlicher-gesundheitsdienst-hygiene-und-infektionsschutz/infektionsschutz/informationen-zum-coronavirus-sars-cov-2/",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: ".links > .textpic-content > .small-12 > p",
				Match:    `insgesamt ([.\d]+).bestätigte`,
			},
			Timestamp: position{
				Selector: ".links > .textpic-content > .small-12 > p",
				Match:    "2.1. 15.04 Uhr",
			},
		},
		"Saarland": {
			URL:       "https://www.saarland.de/SID-C29AF463-5CFEAE1B/253741.htm",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: ".textchapter_frame > p",
				Match:    `landesweit auf ([.\d]+) `,
			},
			Timestamp: position{
				Selector: ".textchapter_frame > p > strong",
				Match:    "2.01.2006 (15:04)",
			},
		},
		"Sachsen": {
			URL:       "https://www.sms.sachsen.de/coronavirus.html",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: "table > tbody > tr:last-child > td:nth-child(2) > strong",
				Match:    `\(([.\d]+)\)`,
			},
			Timestamp: position{
				Selector: ".text-col > p",
				Match:    "Stand: 2. January 2006, 15:04 Uhr",
			},
		},
		"Sachsen-Anhalt": {
			URL: "https://ms.sachsen-anhalt.de/presse/pressemitteilungen/?no_cache=1",
			Listentry: position{
				"#oldpdb > div:first-child > h2 > a",
				"",
			},
			Casecount: position{
				Selector: "table > tbody > tr:last-child > td:last-child",
				Match:    `([.\d]+)`,
			},
			Timestamp: position{
				Selector: ".pm_datum",
				Match:    "den 2. January 2006",
			},
		},
		"Schleswig-Holstein": {
			URL:       "https://www.schleswig-holstein.de/DE/Landesregierung/I/Presse/_documents/Corona-Liste_Kreise.html",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: ".singleview > div > table > tbody > tr:last-child > td:last-child",
				Match:    `([.\d]+)`,
			},
			Timestamp: position{
				Selector: ".singleview > div > table > thead > tr > th:last-child",
				Match:    "Stand 2.01.",
			},
		},
		"Thüringen": {
			URL:       "https://www.landesregierung-thueringen.de/corona-bulletin",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: "table:first-of-type > tbody > tr:nth-child(2) > td:last-child",
				Match:    `([.\d]+)`,
			},
			Timestamp: position{
				Selector: "h3",
				Match:    "Stand: 2. January 2006, 15 Uhr",
			},
		},
	}
}
