package calculator

import "fmt"

type Operation struct {
	Expression string
	Operator   string
	Operands   []float64
}

type Calculator struct {
	Adder Adder
}

func NewCalculator(adder Adder) *Calculator {
	return &Calculator{Adder: adder}
}

func (calculator Calculator) PrintAdd(x, y float64) {
	fmt.Println("Result: ", calculator.Adder.Add(x, y))
}
