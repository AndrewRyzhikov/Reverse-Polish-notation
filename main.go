package main

import (
	"fmt"
	"github.com/andreipimenov/golang-training-2021/03_map_interface_method/homework/notation"
	"github.com/andreipimenov/golang-training-2021/03_map_interface_method/homework/stack"
	"strconv"
)

var operations = map[string]func(a float64, b float64) float64{
	"+": func(a float64, b float64) float64 { return a + b },
	"-": func(a float64, b float64) float64 { return a - b },
	"*": func(a float64, b float64) float64 { return a * b },
	"/": func(a float64, b float64) float64 { return a / b },
}

type Calculator struct{}

type Calc interface {
	Calculate(expression string) float64
}

func (c Calculator) Calculate(expression string) float64 {
	postfix := notation.FromInfixToPostfix(expression)

	var stack stack.FloatStack

	for _, v := range postfix {
		if num, err := strconv.Atoi(v); err == nil {
			stack.Push(float64(num))
		} else if num, err := strconv.ParseFloat(v, 8); err == nil {
			stack.Push(num)
		} else {
			if v != " " {
				a, e1 := stack.Pop()
				b, e2 := stack.Pop()

				if e1 == true && e2 == true {
					val := operations[v](b, a)
					stack.Push(val)
				}
			}
		}
	}

	pop, _ := stack.Pop()

	return pop
}

func main() {
	expression := "20/2-(2+2*3)" // 2
	fmt.Print(Calculator{}.Calculate(expression))
}
