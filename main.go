package main

import (
	"fmt"
	"strings"
)

func main() {
	sum := 0
	fmt.Printf("%-22s %8s    %s\n", "Bundesland", "Fälle", "URL")
	fmt.Println(strings.Repeat("=", 140))
	for name, region := range bundeslaender() {
		casenumber := loadRegion(region)
		if casenumber == 0 {
			fmt.Printf("%-22s %8s    %s\n", name, "n/a", region.URL[8:])
		} else {
			fmt.Printf("%-22s %8d    %s\n", name, casenumber, region.URL[8:])
		}
		sum += casenumber
	}
	fmt.Printf("%-22s %8d    %s\n", "Deutschland", sum, "github.com/HoffmannP/coronaZahlen")
}
