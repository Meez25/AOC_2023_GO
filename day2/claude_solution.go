package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	ID   int
	Sets []map[string]int
}

func parseGame(line string) (Game, error) {
	parts := strings.SplitN(line, ": ", 2)
	if len(parts) != 2 {
		return Game{}, fmt.Errorf("invalid game format: %s", line)
	}

	id, err := strconv.Atoi(strings.TrimPrefix(parts[0], "Game "))
	if err != nil {
		return Game{}, fmt.Errorf("invalid game ID: %s", parts[0])
	}

	sets := []map[string]int{}
	for _, set := range strings.Split(parts[1], "; ") {
		cubes := make(map[string]int)
		for _, cube := range strings.Split(set, ", ") {
			cubeParts := strings.Split(cube, " ")
			if len(cubeParts) != 2 {
				return Game{}, fmt.Errorf("invalid cube format: %s", cube)
			}
			count, err := strconv.Atoi(cubeParts[0])
			if err != nil {
				return Game{}, fmt.Errorf("invalid cube count: %s", cubeParts[0])
			}
			cubes[cubeParts[1]] = count
		}
		sets = append(sets, cubes)
	}

	return Game{ID: id, Sets: sets}, nil
}

func isGamePossible(game Game, limits map[string]int) bool {
	for _, set := range game.Sets {
		for color, count := range set {
			if count > limits[color] {
				return false
			}
		}
	}
	return true
}

func minCubesPower(game Game) int {
	minCubes := make(map[string]int)
	for _, set := range game.Sets {
		for color, count := range set {
			if count > minCubes[color] {
				minCubes[color] = count
			}
		}
	}

	power := 1
	for _, count := range minCubes {
		power *= count
	}
	return power
}

func Claude() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	games := []Game{}
	for scanner.Scan() {
		game, err := parseGame(scanner.Text())
		if err != nil {
			fmt.Printf("Error parsing game: %v\n", err)
			continue // Skip this game and continue with the next
		}
		games = append(games, game)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	limits := map[string]int{"red": 12, "green": 13, "blue": 14}

	part1 := 0
	part2 := 0
	for _, game := range games {
		if isGamePossible(game, limits) {
			part1 += game.ID
		}
		part2 += minCubesPower(game)
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
