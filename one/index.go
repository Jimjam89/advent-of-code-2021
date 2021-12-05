package one

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func loadFile() *os.File {
	file, err := os.Open("one/input.txt")

	if err != nil {
		log.Panic(err)
	}

	return file
}

func Part1() {
	file := loadFile()
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	previous, _ := strconv.Atoi(scanner.Text())
	var current int
	var count int

	for scanner.Scan() {
		current, _ = strconv.Atoi(scanner.Text())

		if current > previous {
			count++
		}

		previous = current
	}

	fmt.Printf("Part 1: %d\n", count)
}

func Part2() {
	var count int
	file := loadFile()
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var values []int
	var current int
	var previous int

	for scanner.Scan() {
		intValue, _ := strconv.Atoi(scanner.Text())
		values = append(values, intValue)
	}

	max := len(values) - len(values)%3

	for i := 0; i < max; i++ {
		current = values[i] + values[i+1] + values[i+2]

		if current > previous && previous != 0 {
			count++
		}

		previous = current
	}

	fmt.Printf("Part 2: %d\n", count)
}
