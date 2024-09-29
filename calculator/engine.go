package calculator

import (
	"fmt"

	"github.com/halfnibble/learn-tdd-go/format"
)

type Adder interface {
	Add(x, y float64) float64
}

type Engine struct {
	expectedLength  int
	validOperations map[string]func(x, y float64) float64
}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) Add(x, y float64) float64 {
	return x + y
}

// GetNumOperands returns the expected number of operands that the engine can process.
func (e *Engine) GetNumOperands() int {
	return e.expectedLength
}

// GetValidOperators returns a slice of the valid operations that the engine accepts.
func (e *Engine) GetValidOperators() []string {
	var ops []string
	for o := range e.validOperations {
		ops = append(ops, o)
	}

	return ops
}

func (e *Engine) ProcessOperation(operation Operation) (*string, error) {
	f, ok := e.validOperations[operation.Operator]
	if !ok {
		err := fmt.Errorf("no operation for operator %s found", operation.Operator)
		return nil, format.Error(operation.Expression, err)
	}
	res := f(operation.Operands[0], operation.Operands[1])
	fres := format.Result(operation.Expression, res)
	return &fres, nil
}
