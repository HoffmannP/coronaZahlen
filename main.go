package main

import (
	"log"

	"time"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type regionType struct {
	name string
	casecount int
	timestamp time.Time
	err error
}

func getRegion(n string, r caseRegion, i chan<- regionType) {
	c, ts, err := r.loadRegion()
	i <- regionType{
		n,
		c,
		ts,
		err,
	}
}

	/*
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
*/

func main() {
	mopo := loadMopo()
	rki, err := loadRKI()
	if err != nil {
		log.Printf("%-22s %s\n", "RKI", err.Error())
	}
	cj, err := loadCj()
	if err != nil {
		log.Printf("%-22s %s\n", "corona.jetzt", err.Error())
	}
	sum := 0
	data := newData("coronaZahlen.json", rki, mopo, cj)
	allRegions := regions()
	next := make(chan regionType)

	for regionName, regionData := range allRegions {
		go getRegion(regionName, regionData, next)
	}
	for i := len(allRegions); i > 0; i-- {
		i := <-next
		if i.err != nil {
			log.Printf("%-22s %s\n", i.name, i.err.Error())
		}
		rkiC := rki.lookup(i.name)
		mopoC := mopo.lookup(i.name)
		cojeC := cj.lookup(i.name)
		sum += max(max(i.casecount, cojeC), max(rkiC, mopoC))
		data.append(i.name, allRegions[i.name].URL, i.timestamp, i.casecount, rkiC, mopoC, cojeC)
	}

	data.sum(sum)
	data.saveJSON("ichplatz/coronaZahlen.json")
}
