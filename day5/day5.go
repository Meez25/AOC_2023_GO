package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type seedRange struct {
	start, end int
}

func Day5() {
	content, err := os.ReadFile("./day5/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	parts := strings.Split(string(content), "\n\n")
	if len(parts) < 2 {
		fmt.Println("Invalid input format")
		return
	}

	// Parse seeds
	seedsInput := strings.Fields(strings.SplitN(parts[0], ":", 2)[1])
	var seeds []seedRange
	for i := 0; i < len(seedsInput); i += 2 {
		if i+1 >= len(seedsInput) {
			fmt.Println("Invalid seed input")
			return
		}
		start, err1 := strconv.Atoi(seedsInput[i])
		length, err2 := strconv.Atoi(seedsInput[i+1])
		if err1 != nil || err2 != nil {
			fmt.Println("Invalid seed numbers")
			return
		}
		seeds = append(seeds, seedRange{start, start + length})
	}

	// Process each mapping block
	for _, block := range parts[1:] {
		var ranges [][]int
		lines := strings.Split(block, "\n")
		for _, line := range lines[1:] { // Skip the header line
			if line == "" {
				continue // Skip empty lines
			}
			var r []int
			for _, num := range strings.Fields(line) {
				n, err := strconv.Atoi(num)
				if err != nil {
					fmt.Printf("Invalid number in mapping: %s\n", num)
					return
				}
				r = append(r, n)
			}
			if len(r) != 3 {
				fmt.Printf("Invalid mapping format: %v\n", r)
				return
			}
			ranges = append(ranges, r)
		}

		var new []seedRange
		for len(seeds) > 0 {
			s, e := seeds[0].start, seeds[0].end
			seeds = seeds[1:]
			mapped := false

			for _, r := range ranges {
				a, b, c := r[0], r[1], r[2]
				os := max(s, b)
				oe := min(e, b+c)
				if os < oe {
					new = append(new, seedRange{os - b + a, oe - b + a})
					if os > s {
						seeds = append(seeds, seedRange{s, os})
					}
					if e > oe {
						seeds = append(seeds, seedRange{oe, e})
					}
					mapped = true
					break
				}
			}

			if !mapped {
				new = append(new, seedRange{s, e})
			}
		}
		seeds = new
	}

	if len(seeds) == 0 {
		fmt.Println("No valid seed ranges found")
		return
	}

	minLocation := seeds[0].start
	for _, sr := range seeds {
		if sr.start < minLocation {
			minLocation = sr.start
		}
	}

	fmt.Println("Minimum location:", minLocation)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
