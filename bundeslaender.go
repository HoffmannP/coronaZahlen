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
				Selector: "figcaption",
				Match:    "Stand: 2. January 2006, 15:04 Uhr",
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
			URL:       "https://msgiv.brandenburg.de/msgiv/de/presse/pressemitteilungen/",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: "h2.bb-teaser-headline",
				// Match:    `Insgesamt ([.\d]+) Erkrankungen
				Match: ` ([.\d]+) bestätigte`,
			},
			Timestamp: position{
				Selector: "h2.bb-teaser-headline + p",
				Match:    " (Stand: 2.01.2006, 15:04 Uhr).",
			},
		},
		"Bremen": {
			URL:       "https://www.gesundheit.bremen.de/sixcms/detail.php?gsid=bremen229.c.32660.de",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: "div.entry-wrapper-normal > table > tbody > tr:nth-child(2) > td:nth-child(2)",
				Match:    `([.\d]+)`,
			},
			Timestamp: position{
				Selector: "div.entry-wrapper-normal > h2",
				Match:    "Monday, 2. January",
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
			URL:       "https://soziales.hessen.de/gesundheit/infektionsschutz/coronavirus-sars-cov-2/taegliche-uebersicht-der-bestaetigten-sars-cov-2-faelle-hessen",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: "table > tbody > tr:last-child > td:nth-child(2)",
				Match:    `([.\d]+)`,
			},
			Timestamp: position{
				Selector: "h3:first-of-type",
				Match:    "Stand 2. January 2006, 15:04 Uhr",
			},
		},
		"Mecklenburg-Vorpommern": {
			URL: "https://www.regierung-mv.de/Landesregierung/wm/Aktuell/?sa.query=neue+Corona-Infektionen&sa.pressemitteilungen.area=11",
			Listentry: position{
				Selector: ".resultlist > .teaser:nth-of-type(2) > div > h3 > a",
				Match:    "https://www.regierung-mv.de",
			},
			Casecount: position{
				Selector: "table > tbody > tr:last-child > td:nth-child(3)",
				Match:    `([.\d]+)`,
			},
			Timestamp: position{
				Selector: "table > tbody > tr:first-child > td:nth-child(3)",
				Match:    "2.01. 15:04",
			},
		},
		"Niedersachsen": {
			URL:       "https://www.apps.nlga.niedersachsen.de/corona/iframe.php",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: "table > tbody > tr:last-child > td:nth-child(2) > span",
				Match:    `([.\d]+)`,
			},
			Timestamp: position{
				Selector: "body > p > b",
				Match:    "Datenstand: 2.01.2006 15:04 Uhr",
			},
		},
		"Nordrhein-Westfalen": {
			URL:       "https://www.mags.nrw/coronavirus-fallzahlen-nrw",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: "table > tbody > tr:last-child > td:nth-child(2)",
				Match:    `([.\d]+)`,
			},
			Timestamp: position{
				Selector: ".field-item > p",
				Match:    "Aktueller Stand: 2. January 2006, 15.04 Uhr.",
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
				Selector: "table > tbody > tr:last-child > td",
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
			URL:       "https://www.coronavirus.sachsen.de/infektionsfaelle-in-sachsen-4151.html",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: "table > tbody > tr:last-child > td:nth-child(2) > strong",
				Match:    `([.\d]+)`,
			},
			Timestamp: position{
				Selector: ".text-col > p",
				Match:    "Stand: 2. January 2006, 15:04 Uhr",
			},
		},
		"Sachsen-Anhalt": {
			URL:       "https://verbraucherschutz.sachsen-anhalt.de/hygiene/infektionsschutz/infektionskrankheiten/coronavirus/#c234506",
			Listentry: position{"", ""},
			Casecount: position{
				Selector: "table > tbody > tr:last-child > td:last-child",
				Match:    `([.\d]+)`,
			},
			Timestamp: position{
				Selector: ".table-responsive + p",
				Match:    "Stand: 2.01.2006 (15:04 Uhr)",
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
