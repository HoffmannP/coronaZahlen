package main

func getRegion(regionName string, regionData caseRegion, rki rkiType, m mopo, j data, sum chan int) {
	casecount, timestamp := regionData.loadRegion()
	rkicount := rki.lookup(regionName)
	mopocount := m.lookup(regionName)
	var summand int

	if casecount == -1 || rkicount > casecount {
		summand = rkicount
	} else {
		summand = casecount
	}

	j.append(regionName, regionData.URL, timestamp, casecount, rkicount, mopocount)

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
	mopo := loadMopo()
	sum := 0
	data := newData("coronaZahlen.json", rki, mopo)
	allRegions := regions()
	next := make(chan int)

	for regionName, regionData := range allRegions {
		go getRegion(regionName, regionData, rki, mopo, data, next)
	}
	remaing(len(allRegions), &sum, next)

	data.sum(sum)
	data.saveJSON("ichplatz/coronaZahlen.json")

}
