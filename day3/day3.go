package day3

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
)

func Day3() {
	file, err := os.Open("./day3/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var dataMap []string

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("read file line error: %v", err)
		}
		dataMap = append(dataMap, line)
	}
	parseDataMapForStars(dataMap)
}

func parseDataMap(dataMap []string) {
	sum := 0
	for i := range dataMap { // for each row
		var detectedPositionOfDigit []int
		for j := 0; j < len(dataMap[0]); j++ { //for each character in line
			if dataMap[i][j] > 47 && dataMap[i][j] < 58 {
				detectedPositionOfDigit = append(detectedPositionOfDigit, j)
			} else {
				// If detectedPositinOfDigit is not empty, calculate if it must be added, then empty the list
				// If it's empty, do nothing
				if len(detectedPositionOfDigit) != 0 {
					fmt.Println(detectedPositionOfDigit)
					sum = sum + calculateValueToAdd(dataMap, detectedPositionOfDigit, i)
					detectedPositionOfDigit = nil
				}
			}
		}
	}
	fmt.Println(sum)
}

func parseDataMapForStars(dataMap []string) {
	sum := 0
	for i := range dataMap { // for each row
		var detectedPositionOfStar []int
		for j := 0; j < len(dataMap[0]); j++ { //for each character in line
			if dataMap[i][j] == 42 {
				detectedPositionOfStar = append(detectedPositionOfStar, j)
			}
		}
		sumGear := gearCalculator(dataMap, detectedPositionOfStar, i)
		sum = sum + sumGear
	}
	fmt.Println(sum)
}

func gearCalculator(dataMap []string, starPosition []int, y int) int {
	sum := 0

	// Look up
	for _, pos := range starPosition {
		var merged []int
		foundTop := detectGearTop(dataMap, pos, y)
		foundMiddle := detectGearMiddle(dataMap, pos, y)
		foundBottom := detectGearBottom(dataMap, pos, y)

		// if y == 1 {
		// 	fmt.Println("foundTop", foundTop, "foundMiddle", foundMiddle, "foundBottom", foundBottom)
		// }

		merged = slices.Concat(foundTop, foundMiddle, foundBottom)
		if len(merged) == 2 {
			fmt.Println("Found gear with sum: ", merged, "line", y, "position", pos)
			sum = sum + merged[0]*merged[1]
		}
	}

	return sum
}

func detectGearTop(dataMap []string, pos int, y int) []int {
	var numberDetectedList []int
	xLeft := pos - 1
	xRight := pos + 2
	// if y == 1 {
	// 	fmt.Println(pos)
	// }
	// if y == 1 && pos == 63 {
	// 	fmt.Println("I should find something here")
	// 	fmt.Println(xLeft, xRight)
	// 	fmt.Println(dataMap[y-1][xLeft:xRight])
	// }
	for i := xLeft; i < xRight; i++ {
		if dataMap[y-1][i] > 47 && dataMap[y-1][i] < 58 {
			numberDetected := extractNumber(dataMap, y-1, i)
			if !slices.Contains(numberDetectedList, numberDetected) {
				numberDetectedList = append(numberDetectedList, numberDetected)
			}
		}
	}
	return numberDetectedList
}

func detectGearMiddle(dataMap []string, pos int, y int) []int {
	var numberDetectedList []int
	xLeft := pos - 1
	xRight := pos + 1
	if dataMap[y][xLeft] > 47 && dataMap[y][xLeft] < 58 {
		numberDetected := extractNumber(dataMap, y, xLeft)
		if !slices.Contains(numberDetectedList, numberDetected) {
			numberDetectedList = append(numberDetectedList, numberDetected)
		}
	}

	if dataMap[y][xRight] > 47 && dataMap[y][xRight] < 58 {
		numberDetected := extractNumber(dataMap, y, xRight)
		if !slices.Contains(numberDetectedList, numberDetected) {
			numberDetectedList = append(numberDetectedList, numberDetected)
		}
	}

	return numberDetectedList
}

func detectGearBottom(dataMap []string, pos int, y int) []int {
	var numberDetectedList []int
	xLeft := pos - 1
	xRight := pos + 2
	for i := xLeft; i < xRight; i++ {
		if dataMap[y+1][i] > 47 && dataMap[y+1][i] < 58 {
			numberDetected := extractNumber(dataMap, y+1, i)
			if !slices.Contains(numberDetectedList, numberDetected) {
				numberDetectedList = append(numberDetectedList, numberDetected)
			}
		}
	}
	return numberDetectedList
}

func extractNumber(dataMap []string, line int, position int) int {
	limitLeft := -1
	limitRight := -1

	for i := position; i >= 0; i-- {
		if dataMap[line][i] < 48 || dataMap[line][i] > 57 {
			limitLeft = i + 1
			break
		} else {
			limitLeft = 0
		}
	}

	for i := position; i <= 140; i++ {
		if dataMap[line][i] < 48 || dataMap[line][i] > 57 {
			limitRight = i
			break
		}
	}

	slice := dataMap[line][limitLeft:limitRight]
	numberFound, _ := strconv.Atoi(slice)
	return numberFound
}

func calculateValueToAdd(dataMap []string, xPoints []int, y int) int {
	top := checkingTop(dataMap, xPoints, y)
	middle := checkingSameLine(dataMap, xPoints, y)
	bottom := checkingBottom(dataMap, xPoints, y)

	if top > 0 {
		return top
	}
	if middle > 0 {
		return middle
	}
	if bottom > 0 {
		return bottom
	}
	return 0

}

func checkingTop(dataMap []string, xPoints []int, y int) int {
	if y > 1 {
		var leftX int
		var rightX int
		if xPoints[0] == 0 {
			leftX = 0
		} else {
			leftX = xPoints[0] - 1
		}

		if xPoints[len(xPoints)-1] == 139 {
			rightX = 139
		} else {
			rightX = xPoints[len(xPoints)-1] + 1
		}

		if y == 139 {
			fmt.Println(leftX, rightX)
		}
		for i := leftX; i <= rightX; i++ {
			if (dataMap[y-1][i] < 48 || dataMap[y-1][i] > 57) && dataMap[y-1][i] != 46 {
				fmt.Println("There is a symbol on top ! ")
				var numberToCompute []byte
				for _, value := range xPoints {
					numberToCompute = append(numberToCompute, dataMap[y][value])
				}

				computedValue, err := strconv.Atoi(string(numberToCompute))
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(numberToCompute, "sum", computedValue)
				return computedValue
			}
		}
	}
	return 0
}

func checkingSameLine(dataMap []string, xPoints []int, y int) int {
	var leftX int
	var rightX int
	if xPoints[0] == 0 {
		leftX = 0
	} else {
		leftX = xPoints[0] - 1
	}

	if xPoints[len(xPoints)-1] == 139 {
		rightX = 139
	} else {
		rightX = xPoints[len(xPoints)-1] + 1
	}

	for i := leftX; i <= rightX; i++ {
		if (dataMap[y][i] < 48 || dataMap[y][i] > 57) && dataMap[y][i] != 46 {
			fmt.Println("There is a symbol on left or right ")
			var numberToCompute []byte
			for _, value := range xPoints {
				numberToCompute = append(numberToCompute, dataMap[y][value])
			}

			computedValue, err := strconv.Atoi(string(numberToCompute))
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(numberToCompute, "sum", computedValue)
			return computedValue
		}
	}
	return 0
}

func checkingBottom(dataMap []string, xPoints []int, y int) int {
	if y < len(dataMap)-1 {
		var leftX int
		var rightX int
		if xPoints[0] == 0 {
			leftX = 0
		} else {
			leftX = xPoints[0] - 1
		}

		if xPoints[len(xPoints)-1] == 139 {
			rightX = 139
		} else {
			rightX = xPoints[len(xPoints)-1] + 1
		}

		for i := leftX; i <= rightX; i++ {
			if (dataMap[y+1][i] < 48 || dataMap[y+1][i] > 57) && dataMap[y+1][i] != 46 {
				fmt.Println("There is a symbol on bottom ! ")
				var numberToCompute []byte
				for _, value := range xPoints {
					numberToCompute = append(numberToCompute, dataMap[y][value])
				}

				computedValue, err := strconv.Atoi(string(numberToCompute))
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(numberToCompute, "sum", computedValue)
				return computedValue
			}
		}
	}

	return 0
}
