package main

func regions() (r caseRegions) {
	return caseRegions{
		"Baden-Würtemberg": {
			URL:      "https://sozialministerium.baden-wuerttemberg.de/de/gesundheit-pflege/gesundheitsschutz/infektionsschutz-hygiene/informationen-zu-coronavirus/",
			Selector: "figcaption",
			Match:    `([.\d]+) bestätigte Corona-Fälle`,
		},
		"Bayern": {
			URL:      "https://www.lgl.bayern.de/gesundheit/infektionsschutz/infektionskrankheiten_a_z/coronavirus/karte_coronavirus/index.htm",
			Selector: "div.row > div:nth-child(2) > table > tbody > tr:nth-child(9) > td:nth-child(2)",
			Match:    `([.\d]+)`,
		},
		"Berlin": {
			URL:      "https://www.berlin.de/sen/gesundheit/themen/gesundheitsschutz-und-umwelt/infektionsschutz/coronavirus/",
			Selector: "",
			Match:    `([.\d]+)`,
		},
		"Brandenburg": {
			URL:      "https://msgiv.brandenburg.de/msgiv/de/start/",
			Selector: "",
			Match:    `([.\d]+)`,
		},
		"Bremen": {
			URL:      "https://www.gesundheit.bremen.de/sixcms/detail.php?gsid=bremen229.c.32657.de",
			Selector: "#main > div:nth-child(3) > div",
			Match:    ` ([.\d]+) als positiv gemeldete`,
		},
		"Hamburg": {
			URL:      "https://www.hamburg.de/bgv/pressemeldungen/13719254/2020-03-13-bgv-coronavirus-aktuell/",
			Selector: "#pushedContainer > main > div.article > div:nth-child(7) > div > div > div",
			Match:    `insgesamt ([.\d]+) positiv Getestete`,
		},
		"Hessen": {
			URL:      "https://soziales.hessen.de/gesundheit/infektionsschutz/coronavirus-sars-cov-2",
			Selector: "section:nth-child(2) > .block-inner > .blockContent",
			Match:    `insgesamt ([.\d]+) SARS-CoV-2-Fälle`,
		},
		"Mecklenburg-Vorpommern": {
			URL:      "https://www.regierung-mv.de/Landesregierung/wm/Aktuelles--Blickpunkte/Wichtige-Informationen-zum-Corona%E2%80%93Virus",
			Selector: "",
			Match:    `([.\d]+)`,
		},
		"Niedersachsen": {
			URL:      "https://www.niedersachsen.de/Coronavirus",
			Selector: ".maincontent > .content > .complementary > div",
			Match:    ` ([.\d]+)\s+laborbestätigte Covid-19-Fälle`,
		},
		"Nordrhein-Westfalen": {
			URL:      "https://www.mags.nrw/coronavirus-fallzahlen-nrw",
			Selector: "table > tbody > tr:nth-child(54) > td:nth-child(2)",
			Match:    `([.\d]+)`,
		},
		"Rheinland-Pfalz": {
			URL:      "https://msagd.rlp.de/de/service/presse/detail/news/News/detail/information-der-landesregierung-zum-aktuellen-stand-hinsichtlich-des-coronavirus-bundesratsinitiati/",
			Selector: "h6",
			Match:    `insgesamt ([.\d]+) bestätigte SARS-CoV-2 Fälle`,
		},
		"Saarland": {
			URL:      "https://www.saarland.de/SID-C29AF463-5CFEAE1B/253741.htm",
			Selector: ".textchapter_frame",
			Match:    `landesweit auf ([.\d]+) – `,
		},
		"Sachsen": {
			URL:      "https://www.sms.sachsen.de/coronavirus.html",
			Selector: "table > tbody > tr:nth-child(14) > td:nth-child(2)",
			Match:    `([.\d]+)`,
		},
		"Sachsen-Anhalt": {
			URL:      "https://ms.sachsen-anhalt.de/themen/gesundheit/aktuell/coronavirus/",
			Selector: "",
			Match:    `([.\d]+)`,
		},
		"Schleswig-Holstein": {
			URL:      "https://www.schleswig-holstein.de/DE/Landesregierung/VIII/_startseite/Artikel_2020/I/200129_Grippe_Coronavirus.html",
			Selector: "",
			Match:    `([.\d]+)`,
		},
		"Thüringen": {
			URL:      "https://www.tmasgff.de/covid-19",
			Selector: "h3",
			Match:    `([.\d]+) bestätigte Infektionen`,
		},
	}
}
