package main

func main() {
	rki := loadRKI()
	sum := 0
	data := newData("CoronaCountsGermany", rki)

	for regionName, regionData := range regions() {
		casecount, timestamp := regionData.loadRegion()
		rkicount := rki.lookup(regionName)

		if casecount == -1 || rkicount > casecount {
			sum += rkicount
		} else {
			sum += casecount
		}

		data.append(
			regionName,
			regionData.URL, timestamp, casecount, rkicount)
	}

	data.sum(sum)
	data.saveHTML("template.html", "docs/index.html")
	data.saveJSON("docs/")

}
