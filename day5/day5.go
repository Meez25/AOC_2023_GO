package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type inputMap struct {
	name   string
	values []string
}

type computedInputMap struct {
	name   string
	values map[int]int
}

func (c *computedInputMap) getValue(inputValue int) int {
	result, ok := c.values[inputValue]
	if !ok {
		return inputValue
	} else {
		return result
	}
}

func newInputMap() inputMap {
	return inputMap{}
}

func newComputerInputMap() computedInputMap {
	values := make(map[int]int)
	name := "Empty"
	return computedInputMap{name: name, values: values}
}

func Day5() {
	var seeds string
	var rawInstructions []string

	file, err := os.Open("./day5/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if strings.Contains(line, "seeds") {
			seeds = line[7:]
		} else if len(line) == 0 {
		} else {
			rawInstructions = append(rawInstructions, line)
		}
	}

	fmt.Println("Seeds :", seeds)
	parsedInstructions := parseInstructions(rawInstructions)

	result := generateResult(seeds, parsedInstructions)

	fmt.Println(result)

}

func parseInstructions(rawInstructions []string) []inputMap {
	var parsedInstructions []inputMap
	inputMap := newInputMap()

	for i, instructionLine := range rawInstructions {
		asRune := []rune(instructionLine)

		if i == len(rawInstructions)-1 {
			startWithDigit := unicode.IsDigit(asRune[0])
			if !startWithDigit {
				inputMap.name = instructionLine
			} else {
				inputMap.values = append(inputMap.values, instructionLine)
				parsedInstructions = append(parsedInstructions, inputMap)
				inputMap = newInputMap()

			}
			break
		}

		nextLineAsRune := []rune(rawInstructions[i+1])
		nextLineStartWithDigit := unicode.IsDigit(nextLineAsRune[0])
		startWithDigit := unicode.IsDigit(asRune[0])
		if !startWithDigit {
			inputMap.name = instructionLine
		} else {
			inputMap.values = append(inputMap.values, instructionLine)
			if !nextLineStartWithDigit {
				parsedInstructions = append(parsedInstructions, inputMap)
				inputMap = newInputMap()
			}
		}

	}
	return parsedInstructions
}

func generateResult(seeds string, computedMaps []inputMap) int {
	min := -1

	var seedsAsIntList []int

	seedsString := strings.Fields(seeds)

	for _, value := range seedsString {
		asInt, _ := strconv.Atoi(value)
		seedsAsIntList = append(seedsAsIntList, asInt)
	}

	seedToSoil := generateCompleteMap(computedMaps[0])
	soilToFertilizer := generateCompleteMap(computedMaps[1])
	fertizilerToWater := generateCompleteMap(computedMaps[2])
	waterToLight := generateCompleteMap(computedMaps[3])
	lightToTemperature := generateCompleteMap(computedMaps[4])
	temperatureToHumidity := generateCompleteMap(computedMaps[5])
	humidityToLocation := generateCompleteMap(computedMaps[6])

	var pipeline = []computedInputMap{
		seedToSoil,
		soilToFertilizer,
		fertizilerToWater,
		waterToLight,
		lightToTemperature,
		temperatureToHumidity,
		humidityToLocation,
	}

	for _, seed := range seedsAsIntList {
		result := seed
		for _, processor := range pipeline {
			result = processor.getValue(result)
		}
		if min == -1 {
			min = result
		} else {
			min = result
		}
	}

	return min
}

func generateCompleteMap(input inputMap) computedInputMap {
	completedInputMap := newComputerInputMap()
	var inputValues []string
	inputValues = input.values
	completedInputMap.name = input.name
	for _, line := range inputValues {
		sourceRangeStart, _ := strconv.Atoi(strings.Fields(line)[1])
		destinationRangeStart, _ := strconv.Atoi(strings.Fields(line)[0])
		rangeLength, _ := strconv.Atoi(strings.Fields(line)[2])
		for i := 0; i < rangeLength; i++ {
			completedInputMap.values[sourceRangeStart+i] = destinationRangeStart + i
		}
	}

	return completedInputMap
}
