package main

import "log"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getRegion(regionName string, regionData caseRegion, rki rkiType, m mopo, j data, sum chan int) {
	casecount, timestamp, err := regionData.loadRegion()
	if err != nil {
		log.Printf("%-22s %s\n", regionName, err.Error())
	}
	rkicount := rki.lookup(regionName)
	mopocount := m.lookup(regionName)
	summand := max(casecount, max(rkicount, mopocount))
	j.append(regionName, regionData.URL, timestamp, casecount, rkicount, mopocount)
	sum <- summand
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
	for i := len(allRegions); i > 0; i-- {
		sum += <-next
	}

	data.sum(sum)
	data.saveJSON("ichplatz/coronaZahlen.json")
}
