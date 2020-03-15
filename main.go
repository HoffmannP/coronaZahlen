package main

import "strconv"

func main() {
	header := []string{"Bundesland", "Fälle", "RKI", "URL"}
	var rows [][]string

	rki := loadRKI()
	sum := 0
	for regionName, regionData := range updateURLs(regions()) {
		var line []string
		line = append(line, regionName)

		casenumber := loadRegion(regionData)
		rki := rki.lookup(regionName)
		if casenumber == -1 {
			line = append(line, "n/a", strconv.Itoa(rki))
			sum += rki
		} else {
			if casenumber > rki {
				line = append(line, "<strong>"+strconv.Itoa(casenumber)+"</strong>", strconv.Itoa(rki))
				sum += casenumber
			} else {
				line = append(line, strconv.Itoa(casenumber), "<strong>"+strconv.Itoa(rki)+"</strong>")
				sum += rki
			}
		}

		line = append(line, "<a href=\""+regionData.URL+"\">Quelle</a>")
		rows = append(rows, line)
	}

	footer := []string{"Deutschland", strconv.Itoa(sum), strconv.Itoa(rki.lookup("Gesamt")), "<a href=\"" + rki.url + "\">Quelle</a>"}

	display(header, rows, footer, "dist/coronaZahl.html")
}
