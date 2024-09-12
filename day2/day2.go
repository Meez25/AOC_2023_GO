package day2

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
)

func Day2() {
	sum := 0
	computedPower := 0
	file, err := os.Open("./day2/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("read file line error: %v", err)
		}
		inputLine := InputLine{
			inputLine: line,
			maxRed:    0,
			maxBlue:   0,
			maxGreen:  0,
		}
		canBeAdded := inputLine.calculateIfPossible()
		if canBeAdded {
			sum = sum + inputLine.gameID
		}
		power := inputLine.calculatePower()
		fmt.Println("Adding", power, "to", computedPower)
		computedPower = computedPower + power
	}
	fmt.Println(sum)
	fmt.Println(computedPower)
}

type InputLine struct {
	inputLine  string
	maxBlue    int
	maxRed     int
	maxGreen   int
	gameID     int
	isPossible bool
}

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

func (inputLine *InputLine) calculatePower() int {
	return inputLine.maxGreen * inputLine.maxRed * inputLine.maxBlue
}

func (inputLine *InputLine) calculateIfPossible() bool {

	// 12 red cubes, 13 green cubes, and 14 blue cubes
	inputLine.calculateGameId()
	inputLine.calculateBlue()
	inputLine.calculateRed()
	inputLine.calculateGreen()
	if inputLine.maxBlue > 14 {
		return false
	}
	if inputLine.maxGreen > 13 {
		return false
	}
	if inputLine.maxRed > 12 {
		return false
	}
	return true
}

func (input *InputLine) calculateGameId() {
	startOfGameID := 5
	var endOfGameID int
	var byteArray []byte
	for i := range input.inputLine {
		if input.inputLine[i] == 58 {
			endOfGameID = i
		}
	}
	for i := startOfGameID; i < endOfGameID; i++ {
		byteArray = append(byteArray, input.inputLine[i])
	}
	computedValue, err := strconv.Atoi(string(byteArray))
	if err != nil {
		fmt.Println(err)
	}
	input.gameID = computedValue
}

func (input *InputLine) calculateBlue() {
	for i := range input.inputLine {
		if string(input.inputLine[i]) == "b" {
			var positionOfBlueDigit []int
			for j := i - 2; j > 0; j-- {
				if input.inputLine[j] > 47 && input.inputLine[j] < 58 {
					positionOfBlueDigit = append(positionOfBlueDigit, j)
				} else {
					break
				}
			}

			var mapOfChar []byte
			for position := range positionOfBlueDigit {
				mapOfChar = append(mapOfChar, input.inputLine[positionOfBlueDigit[position]])
			}
			slices.Reverse(mapOfChar)

			computedValue, err := strconv.Atoi(string(mapOfChar))
			if err != nil {
				fmt.Println(err)
			}
			if computedValue > input.maxBlue {
				input.maxBlue = computedValue
			}
		}
	}
}

func (input *InputLine) calculateRed() {
	for i := range input.inputLine {
		if string(input.inputLine[i]) == "r" && string(input.inputLine[i+1]) == "e" && string(input.inputLine[i+2]) == "d" {
			var positionOfRedDigit []int
			for j := i - 2; j > 0; j-- {
				if input.inputLine[j] > 47 && input.inputLine[j] < 58 {
					positionOfRedDigit = append(positionOfRedDigit, j)
				} else {
					break
				}
			}

			var mapOfChar []byte
			for position := range positionOfRedDigit {
				mapOfChar = append(mapOfChar, input.inputLine[positionOfRedDigit[position]])
			}
			slices.Reverse(mapOfChar)

			computedValue, err := strconv.Atoi(string(mapOfChar))
			if err != nil {
				fmt.Println(err)
			}
			if computedValue > input.maxRed {
				input.maxRed = computedValue
			}
		}
	}
}

func (input *InputLine) calculateGreen() {
	for i := range input.inputLine {
		if string(input.inputLine[i]) == "g" && string(input.inputLine[i+1]) == "r" && string(input.inputLine[i+2]) == "e" {
			var positionOfGreenDigit []int
			for j := i - 2; j > 0; j-- {
				if input.inputLine[j] > 47 && input.inputLine[j] < 58 {
					positionOfGreenDigit = append(positionOfGreenDigit, j)
				} else {
					break
				}
			}

			var mapOfChar []byte
			for position := range positionOfGreenDigit {
				mapOfChar = append(mapOfChar, input.inputLine[positionOfGreenDigit[position]])
			}
			slices.Reverse(mapOfChar)

			computedValue, err := strconv.Atoi(string(mapOfChar))
			if err != nil {
				fmt.Println(err)
			}
			if computedValue > input.maxGreen {
				input.maxGreen = computedValue
			}
		}
	}
}
