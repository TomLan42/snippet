package main

import "fmt"

func main() {

	fmt.Println("Hello, playground")

	var x int32 = -(1 << 31)

	fmt.Println(x)
	x = x >> 1
	fmt.Println(x)

}
