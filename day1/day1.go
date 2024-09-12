package day1

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var digitWords = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func Day1() {
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		fmt.Println("The file could not be read")
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	sum := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("read file line error: %v", err)
		}
		line = strings.TrimSpace(line)
		first := findFirstDigit(line)
		last := findLastDigit(line)
		fmt.Printf("%v, %v%v\n", line, first, last)
		digitToAdd := first*10 + last
		sum += digitToAdd
	}
	fmt.Println(sum)
}

func findFirstDigit(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			return int(s[i] - '0')
		}
		for digit, word := range digitWords {
			if strings.HasPrefix(s[i:], word) {
				return digit + 1
			}
		}
	}
	return 0
}

func findLastDigit(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			return int(s[i] - '0')
		}
		for digit, word := range digitWords {
			if strings.HasSuffix(s[:i+1], word) {
				return digit + 1
			}
		}
	}
	return 0
}
