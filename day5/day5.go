package day5

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type seedsRange struct {
	start int
	size  int
}

type mapRange struct {
	start       int
	destination int
	size        int
}

type mapRangeBlock struct {
	block []mapRange
}

func Day5() {
	file, err := os.ReadFile("./day5/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	seeds, values, _ := bytes.Cut(file, []byte("\n"))
	seedsAsStringList := strings.Fields(strings.Split(string(seeds), ":")[1])
	var seedsAsIntList []seedsRange

	for i := 0; i < len(seedsAsStringList); i = i + 2 {
		start, _ := strconv.Atoi(seedsAsStringList[i])
		end, _ := strconv.Atoi(seedsAsStringList[i+1])
		seedsAsIntList = append(seedsAsIntList, seedsRange{start, end})
	}

	table := bytes.Split(values, []byte("\n\n"))

	var blockOfMap []mapRangeBlock

	for i := 0; i < len(table); i++ {
		var formattedMap []mapRange
		if table[i][0] == 10 {
			table[i][0] = 0
		}
		if table[i][len(table[i])-1] == 10 {
			table[i][len(table[i])-1] = 0
		}
		block := bytes.Split(table[i], []byte("\n"))
		for i := 1; i < len(block); i++ {
			line := strings.Fields(string(block[i]))
			start, _ := strconv.Atoi(line[1])
			size, _ := strconv.Atoi(line[2])
			destination, _ := strconv.Atoi(line[0])
			formattedMap = append(formattedMap, mapRange{destination, start, size})
		}
		blockOfMap = append(blockOfMap, mapRangeBlock{formattedMap})
	}
	// fmt.Println(seedsAsIntList)
	// fmt.Println(blockOfMap)

	for _, seed := range seedsAsIntList {
		var newSeeds []seedsRange
		for _, formattedMap := range blockOfMap {
			fmt.Println("Parse each seed in each block")
			for _, mapRange := range formattedMap.block {
				originalSeed := seed
				var max int
				var min int
				// If is range, do the translation
				if mapRange.start < originalSeed.size+originalSeed.start {
					fmt.Println("Intersection", "seed :", originalSeed.start, originalSeed.start+originalSeed.size, "range", mapRange.start, mapRange.destination)
					min = mapRange.start
				}
				if originalSeed.start+originalSeed.size < mapRange.destination {
					fmt.Println("Intersection", "seed :", originalSeed.start, originalSeed.start+originalSeed.size, "range", mapRange.start, mapRange.destination)
					max = mapRange.destination
				} else {
					min = seed.start
					max = seed.size
					fmt.Println("Sortir l'input tel quel")
					newSeeds = append(newSeeds, seedsRange{min, max})
				}

				fmt.Println("Added seed", seedsRange{min, max})

				// If not in range, do nothing
			}
		}
		fmt.Println(newSeeds)
		seedsAsIntList = newSeeds
	}

}
