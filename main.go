package main

import (
	"sort"
	"strconv"
)

func main() {
	rki := loadRKI()
	sum := 0

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
				line = append(line, strconv.Itoa(casecount), "<span class=\"blass\">"+strconv.Itoa(rki)+"</span>")
				sum += casecount
			case rki > casecount:
				line = append(line, "<span class=\"blass\">"+strconv.Itoa(casecount)+"</span>", strconv.Itoa(rki))
				sum += casecount
			default:
				line = append(line, strconv.Itoa(casecount), strconv.Itoa(rki))
				sum += casecount
			}
		}

		line = append(line, timestamp.Format("2.01.2006 15:04 Uhr"))
		rows = append(rows, line)
	}

	footer := []string{
		"Deutschland",
		strconv.Itoa(sum),
		strconv.Itoa(rki.lookup("Gesamt")),
		rki.timestamp.Format("RKI: 2.01.2006 15:04 Uhr "),
	}

	display(header, rows, footer, "docs/index.html")
}
