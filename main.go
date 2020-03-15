package main

import (
	"sort"
	"strconv"
)

func main() {
	header := []string{"Bundesland", "Fälle", "RKI", "Stand"}
	var rows [][]string

	rki := loadRKI()
	sum := 0

	regions := regions()
	var regionNames []string
	for regionName := range regions {
		regionNames = append(regionNames, regionName)
	}
	sort.Strings(regionNames)

	for _, regionName := range regionNames {
		regionData := regions[regionName]
		var line []string
		line = append(line, regionName)

		casecount, timestamp := loadRegion(regionData)
		rki := rki.lookup(regionName)
		if casecount == -1 {
			line = append(line, "n/a", strconv.Itoa(rki))
			sum += rki
		} else {
			if casecount > rki {
				line = append(line, "<strong>"+strconv.Itoa(casecount)+"</strong>", strconv.Itoa(rki))
				sum += casecount
			} else {
				line = append(line, strconv.Itoa(casecount), "<strong>"+strconv.Itoa(rki)+"</strong>")
				sum += rki
			}
		}

		line = append(line, "<a href=\""+regionData.URL+"\">"+timestamp.Format("2.01.2006 15:04")+"</a>")
		rows = append(rows, line)
	}

	footer := []string{
		"Deutschland",
		strconv.Itoa(sum),
		strconv.Itoa(rki.lookup("Gesamt")),
		"<a href=\"" + rki.timestamp.Format("2.01.2006 15:04") + "\">Quelle</a>",
	}

	display(header, rows, footer, "docs/index.html")
}
