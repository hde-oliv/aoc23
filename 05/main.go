package main

import (
	"fmt"
	"math"
	"slices"
)

var seeds = []uint64{}

var seedToSoil = [][]uint64{}

var soilToFertilizer = [][]uint64{}

var fertilizerToWater = [][]uint64{}

var waterToLight = [][]uint64{}

var lightToTemperature = [][]uint64{}

var temperatureToHumidity = [][]uint64{}

var humidityToLocation = [][]uint64{}

func findCorrespondence(ls [][]uint64, seed uint64) uint64 {
	for _, v := range ls {
		if seed >= v[1] && seed <= (v[1]+v[2]-1) {
			diff := seed - v[1]
			return v[0] + diff
		}
	}
	return seed
}

func main() {
	l := make([]uint64, 0)

	// Part 01
	for _, s := range seeds {
		s = findCorrespondence(seedToSoil, s)
		s = findCorrespondence(soilToFertilizer, s)
		s = findCorrespondence(fertilizerToWater, s)
		s = findCorrespondence(waterToLight, s)
		s = findCorrespondence(lightToTemperature, s)
		s = findCorrespondence(temperatureToHumidity, s)
		s = findCorrespondence(humidityToLocation, s)
		l = append(l, s)
	}

	slices.Sort(l)
	fmt.Println(l)
	fmt.Println(l[0])

	// Part 02
	l = nil
	for i := 0; i < len(seeds); i += 2 {
		var min uint64 = math.MaxUint64

		s := seeds[i]
		e := seeds[i+1]

		for e != 0 {
			r := findCorrespondence(seedToSoil, s)
			r = findCorrespondence(soilToFertilizer, r)
			r = findCorrespondence(fertilizerToWater, r)
			r = findCorrespondence(waterToLight, r)
			r = findCorrespondence(lightToTemperature, r)
			r = findCorrespondence(temperatureToHumidity, r)
			r = findCorrespondence(humidityToLocation, r)

			if r < min {
				min = r
			}

			s++
			e--
		}

		l = append(l, min)
	}

	slices.Sort(l)
	fmt.Println(l)
	fmt.Println(l[0])
}
