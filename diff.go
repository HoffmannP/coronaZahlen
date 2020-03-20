package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func getOldRegions(filename string) (c casecounts, err error) {
	j, err := ioutil.ReadFile(filename)
	if err != nil {
		return c, err
	}

	var d data
	if err := json.Unmarshal(j, &d); err != nil {
		return c, err
	}
	return d.Regions, nil
}

func diff(filename string, regions *casecounts) bool {
	oldRegions, err := getOldRegions(filename)
	if err != nil {
		log.Println("Diff: " + err.Error())
		return true
	}
	for regionName, regionOldData := range oldRegions {
		regionNewData := (*regions)[regionName]
		if regionOldData.Count != regionNewData.Count {
			if regionNewData.Count > 0 {
				return true
			}
			regionNewData.Count = regionOldData.Count
			regionNewData.Date = regionOldData.Date
			(*regions)[regionName] = regionNewData
		}
		if (regionOldData.Date != regionNewData.Date) ||
			(regionOldData.RKI != regionNewData.RKI) ||
			(regionOldData.Mopo != regionNewData.Mopo) ||
			(regionOldData.CJ != regionNewData.CJ) {
			return true
		}
	}
	return false
}
