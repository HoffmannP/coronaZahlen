package main

import (
	"sort"
	"strconv"
	"time"
)

func blass(t string) string {
	return "<span class=\"blass\">" + t + "</span>"
}

func main() {
	rki := loadRKI()
	sum := 0
	json := newJSON("CoronaCountsGermany")

	header := []string{"Bundesland", "Fälle", "<a href=\"" + rki.url + "\">RKI</a>", "Stand"}
	var rows [][]string

	regions := regions()
	var regionNames []string
	for regionName := range regions {
		regionNames = append(regionNames, regionName)
	}
	sort.Strings(regionNames)

	for _, regionName := range regionNames {
		regionData := regions[regionName]
		var line []string

		casecount, timestamp := regionData.loadRegion()
		rki := rki.lookup(regionName)
		line = append(line, "<a href=\""+regionData.URL+"\">"+regionName+"</a>")

		if casecount == -1 {
			line = append(line, "n/a", strconv.Itoa(rki))
			sum += rki
		} else {
			switch {
			case casecount > rki:
				line = append(line, strconv.Itoa(casecount), blass(strconv.Itoa(rki)))
				sum += casecount
			case rki > casecount:
				line = append(line, blass(strconv.Itoa(casecount)), strconv.Itoa(rki))
				sum += casecount
			default:
				line = append(line, strconv.Itoa(casecount), strconv.Itoa(rki))
				sum += casecount
			}
		}

		line = append(line, timestamp.Format("2.01.2006 15:04 Uhr"))
		rows = append(rows, line)
		json.append(regionName, regionData.URL, timestamp, casecount, rki)
	}

	gesamtRki := rki.lookup("Gesamt")
	footer := []string{
		"Deutschland",
		strconv.Itoa(sum),
		strconv.Itoa(gesamtRki),
		rki.timestamp.Format("RKI: 2.01.2006 15:04 Uhr "),
	}
	json.append("Deutschland", "https://hoffis-eck.de/coronaZahlen/"+json.getName()+".json", time.Now(), sum, gesamtRki)

	saveHTML(header, rows, footer, json.getName(), "docs/index.html")
	json.save("docs/")

}
