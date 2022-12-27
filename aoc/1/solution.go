package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

	nums := stringToInt(fileLines)
	fmt.Println(solutionPart1(nums))
	fmt.Println(solutionPart2(nums))

}

func stringToInt(input []string) []int {
	result := make([]int, 0)
	groupSum := 0
	for _, v := range input {
		if v == "" {
			result = append(result, groupSum)
			groupSum = 0
		} else {
			num, _ := strconv.Atoi(v)
			groupSum += num
		}
	}
	return result
}

func solutionPart1(nums []int) int {
	max := 0
	for _, num := range nums {
		if num >= max {
			max = num
		}
	}
	return max
}

// let top3[0] always store the smallest element among top3
func solutionPart2(nums []int) int {
	top3 := make([]int, 3)

	for _, num := range nums {
		if num > top3[0] {
			top3[0] = num
			min := min(top3[1], top3[2])
			if min < top3[0] {
				if min == top3[1] {
					swap(top3, 0, 1)
				}

				if min == top3[2] {
					swap(top3, 0, 2)
				}
			}
		}
	}

	return top3[0] + top3[1] + top3[2]
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
