package main

import (
	"bufio"
	"fmt"
	"os"
)

var scoreMap = map[string]int{
	"A X": 1 + 3,
	"B X": 1 + 0,
	"C X": 1 + 6,
	"A Y": 2 + 6,
	"B Y": 2 + 3,
	"C Y": 2 + 0,
	"A Z": 3 + 0,
	"B Z": 3 + 6,
	"C Z": 3 + 3,
}

var scoreMap2 = map[string]int{
	"A X": 1 + 0,
	"B X": 1 + 0,
	"C X": 1 + 0,
	"A Y": 2 + 3,
	"B Y": 2 + 3,
	"C Y": 2 + 3,
	"A Z": 3 + 6,
	"B Z": 3 + 6,
	"C Z": 3 + 6,
}

func main() {
	filepath := "./input"
	readFile, err := os.Open(filepath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()

	fmt.Println(solutionPart1(fileLines))
	fmt.Println(solutionPart2(fileLines))

}

func solutionPart1(pairs []string) int {
	sum := 0

	for _, pair := range pairs {
		sum += scoreMap[pair]
	}

	return sum
}

func solutionPart2(pairs []string) int {

	sum := 0

	for _, pair := range pairs {
		sum += scoreMap2[pair]
	}

	return sum
}
