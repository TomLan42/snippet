package main

import "fmt"

func main() {
	bitshift()
	initilization()
}

func bitshift() {
	fmt.Println("Hello, playground")

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
