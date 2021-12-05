package two

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	action string
	value  int
}

func loadInstructions() []instruction {
	file, err := os.Open("two/input.txt")

	if err != nil {
		log.Panic(err)
	}

	defer file.Close()

	var instructions []instruction

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		strings := strings.Fields(scanner.Text())
		value, _ := strconv.Atoi(strings[1])
		instructions = append(instructions, instruction{strings[0], value})
	}

	return instructions
}

func Part1() {
	instructions := loadInstructions()
	var horizontalPos int
	var depth int

	for _, instruction := range instructions {
		switch instruction.action {
		case "forward":
			horizontalPos += instruction.value
		case "down":
			depth += instruction.value
		case "up":
			depth -= instruction.value
		}

	}

	fmt.Printf("Part 1: %d\n", horizontalPos*depth)
}

func Part2() {
	instructions := loadInstructions()
	var horizontalPos int
	var depth int
	var aim int

	for _, instruction := range instructions {
		switch instruction.action {
		case "forward":
			horizontalPos += instruction.value
			depth += instruction.value * aim
		case "down":
			aim += instruction.value
		case "up":
			aim -= instruction.value
		}

	}

	fmt.Printf("Part 2: %d\n", horizontalPos*depth)
}
