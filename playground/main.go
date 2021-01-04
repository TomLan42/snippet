package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	bitshift()
	initilization()
	playwithstruct()
	sleeptest()
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
