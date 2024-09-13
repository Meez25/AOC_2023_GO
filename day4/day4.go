package day4

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"strconv"
	"strings"
)

func Day4() {
	file, err := os.Open("./day4/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var fileLines []string

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("read file line error: %v", err)
		}

		fileLines = append(fileLines, line)
	}

	parseFile(fileLines)

	// fileLines, _ = computeFile(fileLines)

}

func getLineNoOfFirstGame(gameNo int, fileLines []string) int {
	for lineNo, line := range fileLines {
		if getGameNo(line) == gameNo {
			return lineNo
		}
	}
	return 0
}

func parseFile(fileLines []string) {
	currentDeck := make(map[int]int)
	for i := 0; i < len(fileLines); i++ {
		currentDeck[i+1] = 1
	}

	for _, line := range fileLines {
		numberOfCardToAdd := parseLine(line)
		gameNo := getGameNo(line)

		for j := 0; j < numberOfCardToAdd; j++ {
			for k := 0; k < currentDeck[gameNo]; k++ {
				currentDeck[gameNo+j+1] = currentDeck[gameNo+j+1] + 1
			}
		}

	}

	sum := 0
	for i := range currentDeck {
		sum = sum + currentDeck[i]
	}
	fmt.Println(sum)
}

func getGameNo(line string) int {
	gameNo := strings.Fields(line[0:8])[1]
	gameNoAsInt, _ := strconv.Atoi(gameNo)
	return gameNoAsInt
}

func parseLine(line string) int {
	sum := 0
	winningNumbersString := strings.Fields(line[10:39])
	playerNumbers := strings.Fields(line[41:])
	for _, winningNumber := range winningNumbersString {
		for _, playerNumber := range playerNumbers {
			if winningNumber == playerNumber {
				sum = sum + 1
			}
		}
	}
	return sum
}
