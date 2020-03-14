package main

import (
	"fmt"
	"strings"
)

func main() {
	rki := loadRKI()
	sum := 0
	fmt.Printf("%-22s %8s %8s    %s\n", "Bundesland", "Fälle", "RKI", "URL")
	fmt.Println(strings.Repeat("=", 140))
	for regionName, regionData := range regions() {
		casenumber := loadRegion(regionData)
		rki := rki.lookup(regionName)
		if casenumber == 0 {
			fmt.Printf("%-22s %8s %8d    %s\n", regionName, "n/a", rki, regionData.URL[8:])
		} else {
			fmt.Printf("%-22s %8d %8d    %s\n", regionName, casenumber, rki, regionData.URL[8:])
		}
		if rki > casenumber {
			sum += rki
		} else {
			sum += casenumber
		}
	}
	fmt.Printf("%-22s %8d %8d    %s\n", "Deutschland", sum, rki.lookup("Gesamt"), "github.com/HoffmannP/coronaZahlen")
}
