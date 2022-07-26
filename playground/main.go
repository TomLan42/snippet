package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// bitshift()
	// initilization()
	// playwithstruct()
	// sleeptest()
	// appendtest()

	fmt.Println(partition("ababbbabbaba"))
}

func bitshift() {
	fmt.Println("Hello, playgrdound")

	var x int32 = -(1 << 31)

	fmt.Println(x)
	x = x >> 1
	fmt.Println(x)
}

func initilization() {

	var x map[int]int

	fmt.Println(x)

	y := make(map[int]int)
	fmt.Println(y[0])
	fmt.Println(x[2])
}

func playwithstruct() {
	type S struct {
		name string
	}
	A := &S{
		name: "HHHH"}
	A.name = "GOUGOU"
	fmt.Println(A.name)
	B := S{
		name: "WHYWHY"}
	B.name = "YESYES"
	fmt.Println(B.name)
}

func sleeptest() {
	var x int = 1
	var speed float64 = 12
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(x) * time.Duration(float64(time.Second)/speed))
		log.Println(i)
	}
}

func appendtest() {

	var result [][]int
	var cur []int = []int{1, 2}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		helperAppend(append(cur, i), &result, i)
		// fmt.Println(cur)
	}

}

func helperAppend(x []int, result *[][]int, count int) {

	if len(*result) == 4 {
		return
	}
	fmt.Println(x)
	*result = append(*result, x)
	fmt.Println(result)
}

func partition(s string) [][]string {

	result := [][]string{}
	if len(s) == 0 {
		return result
	}
	cur := []string{}

	dfs(s, 0, cur, &result)

	return result

}

func dfs(s string, idx int, cur []string, result *[][]string) {
	start, end := idx, len(s)

	if start == end {

		*result = append(*result, cur)

		return
	}

	for i := start; i < end; i++ {
		if isPal(s, start, i) {

			dfs(s, i+1, append(cur, s[start:i+1]), result)
		}
	}

}

func isPal(str string, s, e int) bool {

	for s < e {
		if str[s] != str[e] {
			return false
		}
		s++
		e--
	}
	return true
}
