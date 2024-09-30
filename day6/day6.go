package day6

import "fmt"

func Day6() {
	// race1 := computeRace(7, 9)
	// race2 := computeRace(15, 40)
	// race3 := computeRace(30, 200)
	race1 := computeRace(40817772, 219101213651089)

	fmt.Print(race1)
}

func computeRace(time, distance int) int {
	table := make([]int, time)
	count := 0

	for i := 0; i < time; i++ {
		if i == 0 {
			table[i] = 0
		}
		table[i] = i * (time - i)
		if table[i] > distance {
			count++
		}
	}
	return count
}
