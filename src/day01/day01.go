package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

func checkError(e error) {

	if e != nil {
		panic(e)
	}
}

func fuelCalculation(mass int) int {

	return (mass / 3) - 2
}

func part1(lines []string) {

	totalFuel := 0

	for _, value := range lines {

		//2. Rstrip
		value = strings.TrimRightFunc(value, unicode.IsSpace)

		// 3. Convert to int
		distance, err := strconv.Atoi(value)
		checkError(err)

		// 4. Calculation
		totalFuel += fuelCalculation(distance)

	}

	fmt.Println("Part 1 The total fuel required : ", totalFuel)

}

func part2(lines []string) {

	totalFuel := 0

	for _, value := range lines {

		//2. Rstrip
		value = strings.TrimRightFunc(value, unicode.IsSpace)

		// 3. Convert to int
		distance, err := strconv.Atoi(value)
		checkError(err)

		// 4. Calculation
		fuelNeeded := fuelCalculation(distance)
		totalFuel += fuelNeeded

		for fuelNeeded > 6 {

			fuelNeeded = fuelCalculation(fuelNeeded)
			totalFuel += fuelNeeded
		}

	}

	fmt.Println("Part 2 The total fuel required + the fuel required for the fuel : ", totalFuel)

}

func main() {

	// This is how you dump all the contents of the file to memory
	// The file that is returned is []byte
	data, err := ioutil.ReadFile("input.txt")
	checkError(err)

	// 1. Convert byte data to string data, and split it to new lines
	lines := strings.Split(string(data), "\n")

	part1(lines)
	part2(lines)
}
