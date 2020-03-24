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
	name      string
	casecount int
	timestamp time.Time
	err       error
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
	waitingFor := 0

	for regionName, regionData := range allRegions {
		/*
			if regionName != "Bremen" {
				continue
			}
		*/
		go getRegion(regionName, regionData, next)
		waitingFor++
	}
	for i := waitingFor; i > 0; i-- {
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
