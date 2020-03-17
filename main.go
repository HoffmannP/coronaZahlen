package main

func getRegion(regionName string, regionData caseRegion, rki rkiType, j data, sum chan int) {
	casecount, timestamp := regionData.loadRegion()
	rkicount := rki.lookup(regionName)
	var summand int

	if casecount == -1 || rkicount > casecount {
		summand = rkicount
	} else {
		summand = casecount
	}

	j.append(regionName, regionData.URL, timestamp, casecount, rkicount)

	sum <- summand
}

func remaing(remaining int, sum *int, next chan int) {
	for remaining > 0 {
		*sum += <-next
		remaining--
	}
}

func main() {
	rki := loadRKI()
	sum := 0
	data := newData("CoronaCountsGermany", rki)
	allRegions := regions()
	next := make(chan int)

	for regionName, regionData := range allRegions {
		go getRegion(regionName, regionData, rki, data, next)
	}
	remaing(len(allRegions), &sum, next)

	data.sum(sum)
	data.saveHTML("template.html", "docs/index.html")
	data.saveJSON("docs/")

}
