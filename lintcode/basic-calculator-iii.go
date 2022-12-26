package lintcode

import (
	"fmt"
	"strconv"
)

// method 1
// 1.1 Convert infix expression to tree. Then convert tree to reverse polish/or eval the tree.
// 1.2 Convert infix expression to reverse polish using stack

// method 2
// recursive method

type Node struct {
	val   rune
	left  *Node
	right *Node
}

var priorityMap map[string]int = map[string]int{
	"+": 0,
	"-": 0,
	"*": 1,
	"/": 1,
	"(": -1,
}

func InfixToSuffix(tokens []string) []string {
	stack := make([]string, 0)
	result := make([]string, 0)
	for i := 0; i < len(tokens); i++ {
		cur := tokens[i]
		switch cur {
		case "+", "-", "*", "/":
			for len(stack) > 0 && priorityMap[stack[len(stack)-1]] >= priorityMap[cur] {
				fmt.Println("detail: cur is ", cur, stack)
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				result = append(result, top)
			}
			stack = append(stack, cur)
		case "(":
			stack = append(stack, cur)
		case ")":
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if top == ("(") {
					break
				}
				result = append(result, top)
			}
		default:
			result = append(result, cur)
		}
		fmt.Println("cur is ", cur, stack)
	}

	// clear anything left in the stack
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, top)
	}

	return result
}

func Calculate(s string) float64 {
	// tokenize string
	tokens := make([]string, 0)
	for _, rune := range s {
		if rune == ' ' {
			continue
		}

		tokens = append(tokens, string(rune))
	}

	// convert to RPN
	rpnTokens := InfixToSuffix(tokens)

	// compute RPN
	stack := make([]float64, 0)
	for _, token := range rpnTokens {
		switch token {
		case "*", "/", "+", "-":
			rand1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			rand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if token == "+" {
				stack = append(stack, rand1+rand2)
			}

			if token == "-" {
				stack = append(stack, rand2-rand1)
			}

			if token == "*" {
				stack = append(stack, rand1*rand2)
			}

			if token == "/" {
				stack = append(stack, rand2/rand1)
			}
		default:
			digit, _ := strconv.Atoi(token)
			stack = append(stack, float64(digit))
		}
	}

	return stack[len(stack)-1]
}
