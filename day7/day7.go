package day7

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day7() {
	// inputTable := parseInput("day7/input.txt")
	inputTable := parseInput("day7/input_large.txt")
	result := computeInput(inputTable)
	resultWithJoker := computeInputJoker(inputTable)
	fmt.Println(result)
	fmt.Println(resultWithJoker)
}

func computeInput(inputTable []string) int {
	slices.SortFunc(inputTable, pokerCmp)
	sum := 0
	for i, val := range inputTable {
		value := strings.Fields(val)[1]
		asInt, _ := strconv.Atoi(value)
		sum = sum + asInt*(i+1)
	}
	return sum
}

func computeInputJoker(inputTable []string) int {
	slices.SortFunc(inputTable, pokerCmpJoker)
	sum := 0
	for i, val := range inputTable {
		value := strings.Fields(val)[1]
		asInt, _ := strconv.Atoi(value)
		sum = sum + asInt*(i+1)
	}
	return sum
}
func pokerCmpJoker(a, b string) int {
	values := []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}
	hand1 := FindBest(a)
	hand2 := FindBest(b)
	if hand1 < hand2 {
		return -1
	} else if hand2 < hand1 {
		return 1
	} else {
		// Equal
		for i := 0; i < 5; i++ {
			strengthA := slices.Index(values, string(a[i]))
			strengthB := slices.Index(values, string(b[i]))
			if strengthA > strengthB {
				return 1
			} else if strengthB > strengthA {
				return -1
			} else {
				continue
			}
		}
		return 0
	}
}

func pokerCmp(a, b string) int {
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	hand1 := findHand(a)
	hand2 := findHand(b)
	if hand1 < hand2 {
		return -1
	} else if hand2 < hand1 {
		return 1
	} else {
		// Equal
		for i := 0; i < 5; i++ {
			strengthA := slices.Index(values, string(a[i]))
			strengthB := slices.Index(values, string(b[i]))
			if strengthA > strengthB {
				return 1
			} else if strengthB > strengthA {
				return -1
			} else {
				continue
			}
		}
		return 0
	}
}

func FindBest(input string) int {
	input = input[:5]
	max := 0
	var uniqueChar []rune
	for _, char := range input {
		if slices.Contains(uniqueChar, char) {
			continue
		} else {
			uniqueChar = append(uniqueChar, char)
		}
	}
	// No J
	if strings.Index(input, "J") == -1 {
		maxValue := findHand(input)
		return maxValue
	}
	inputAsTable := strings.Split(input, "")

	var jposition []int
	for i := range inputAsTable {
		if inputAsTable[i] == "J" {
			jposition = append(jposition, i)
		}
	}
	if len(jposition) == 1 {
		for _, uniqueVal := range uniqueChar {
			inputAsTable[jposition[0]] = string(uniqueVal)
			rebuild := strings.Join(inputAsTable, "")
			maxValue := findHand(rebuild)
			if maxValue > max {
				max = maxValue
			}
		}
	}

	if len(jposition) == 2 {
		for _, uniqueVal := range uniqueChar {
			inputAsTable[jposition[0]] = string(uniqueVal)
			for _, uniqueVal := range uniqueChar {
				inputAsTable[jposition[1]] = string(uniqueVal)
				rebuild := strings.Join(inputAsTable, "")
				maxValue := findHand(rebuild)
				if maxValue > max {
					max = maxValue
				}
			}
		}
	}
	if len(jposition) == 3 {
		for _, uniqueVal := range uniqueChar {
			inputAsTable[jposition[0]] = string(uniqueVal)
			for _, uniqueVal := range uniqueChar {
				inputAsTable[jposition[1]] = string(uniqueVal)
				for _, uniqueVal := range uniqueChar {
					inputAsTable[jposition[2]] = string(uniqueVal)
					rebuild := strings.Join(inputAsTable, "")
					maxValue := findHand(rebuild)
					if maxValue > max {
						max = maxValue
					}
				}
			}
		}
	}
	if len(jposition) == 4 {
		for _, uniqueVal := range uniqueChar {
			inputAsTable[jposition[0]] = string(uniqueVal)
			for _, uniqueVal := range uniqueChar {
				inputAsTable[jposition[1]] = string(uniqueVal)
				for _, uniqueVal := range uniqueChar {
					inputAsTable[jposition[2]] = string(uniqueVal)
					for _, uniqueVal := range uniqueChar {
						inputAsTable[jposition[3]] = string(uniqueVal)
						rebuild := strings.Join(inputAsTable, "")
						maxValue := findHand(rebuild)
						if maxValue > max {
							max = maxValue
						}
					}
				}
			}
		}
	}
	if len(jposition) == 5 {
		max = 7
	}
	return max
}

func findHandWithJoker(input string) int {
	parts := strings.Fields(input)

	if isFiveOfAKind(parts[0]) {
		return 7
	}
	if isFourOfAKind(parts[0]) {
		return 6
	}
	if isFullHouse(parts[0]) {
		return 5
	}
	if isThreeOfAKind(parts[0]) {
		return 4
	}
	if isTwoPair(parts[0]) {
		return 3
	}
	if isPair(parts[0]) {
		return 2
	}
	if isHighCard(parts[0]) {
		return 1
	}
	return 0
}

func findHand(input string) int {
	parts := strings.Fields(input)

	if isFiveOfAKind(parts[0]) {
		return 7
	}
	if isFourOfAKind(parts[0]) {
		return 6
	}
	if isFullHouse(parts[0]) {
		return 5
	}
	if isThreeOfAKind(parts[0]) {
		return 4
	}
	if isTwoPair(parts[0]) {
		return 3
	}
	if isPair(parts[0]) {
		return 2
	}
	if isHighCard(parts[0]) {
		return 1
	}
	return 0
}

func isFiveOfAKind(input string) bool {
	if strings.Count(input, string(input[0])) == 5 {
		return true
	} else {
		return false
	}
}

func isFourOfAKind(input string) bool {
	if strings.Count(input, string(input[0])) == 4 || strings.Count(input, string(input[1])) == 4 {
		return true
	} else {
		return false
	}
}

func isFullHouse(input string) bool {
	var uniqueChar []rune
	for _, char := range input {
		if slices.Contains(uniqueChar, char) {
			continue
		} else {
			uniqueChar = append(uniqueChar, char)
		}
	}
	if len(uniqueChar) != 2 {
		return false
	}
	if len(uniqueChar) == 2 && (strings.Count(input, string(uniqueChar[0])) == 2 || strings.Count(input, string(uniqueChar[0])) == 3) {
		return true
	} else {
		return false
	}
}

func isThreeOfAKind(input string) bool {
	var uniqueChar []rune
	for _, char := range input {
		if slices.Contains(uniqueChar, char) {
			continue
		} else {
			uniqueChar = append(uniqueChar, char)
		}
	}
	if len(uniqueChar) != 3 {
		return false
	}
	if len(uniqueChar) == 3 && (strings.Count(input, string(uniqueChar[0])) == 3 || strings.Count(input, string(uniqueChar[1])) == 3) || strings.Count(input, string(uniqueChar[2])) == 3 {
		return true
	} else {
		return false
	}
}

func isTwoPair(input string) bool {
	var uniqueChar []rune
	for _, char := range input {
		if slices.Contains(uniqueChar, char) {
			continue
		} else {
			uniqueChar = append(uniqueChar, char)
		}
	}
	if len(uniqueChar) != 3 {
		return false
	}
	if len(uniqueChar) == 3 && (strings.Count(input, string(uniqueChar[0])) == 2 || strings.Count(input, string(uniqueChar[1])) == 2) || strings.Count(input, string(uniqueChar[2])) == 2 {
		return true
	} else {
		return false
	}
}

func isPair(input string) bool {
	var uniqueChar []rune
	for _, char := range input {
		if slices.Contains(uniqueChar, char) {
			continue
		} else {
			uniqueChar = append(uniqueChar, char)
		}
	}
	if len(uniqueChar) != 4 {
		return false
	}
	if len(uniqueChar) == 4 && (strings.Count(input, string(uniqueChar[0])) == 2 || strings.Count(input, string(uniqueChar[1])) == 2) || strings.Count(input, string(uniqueChar[2])) == 2 || strings.Count(input, string(uniqueChar[3])) == 2 {
		return true
	} else {
		return false
	}
}

func isHighCard(input string) bool {
	var uniqueChar []rune
	for _, char := range input {
		if slices.Contains(uniqueChar, char) {
			continue
		} else {
			uniqueChar = append(uniqueChar, char)
		}
	}
	if len(uniqueChar) != 5 {
		return false
	} else {
		return true
	}
}

func parseInput(fileName string) []string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Could not read file %w", err)
	}
	fileString := string(file)
	inputTable := strings.SplitN(fileString, "\n", -1)
	inputTable = inputTable[:len(inputTable)-1]
	return inputTable
}
