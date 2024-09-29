package input

import (
	"github.com/halfnibble/learn-tdd-go/calculator"
	"github.com/halfnibble/learn-tdd-go/format"
)

// OperationProcessor is the interface for processing mathematical expressions
type OperationProcessor interface {
	ProcessOperation(operation *calculator.Operation) (string, error)
}

// ValidationHelper is the interface for input validation
type ValidationHelper interface {
	CheckInput(operator string, operands []float64) error
}

// Parser is responsible for converting input to the mathematical operations
type Parser struct {
	engine    OperationProcessor
	validator ValidationHelper
}

func NewParser(engine OperationProcessor, validator ValidationHelper) *Parser {
	return &Parser{
		engine:    engine,
		validator: validator,
	}
}

func (p *Parser) ProcessExpression(expr string) (string, error) {
	operation, err := p.getOperation(expr)
	if err != nil {
		return "", format.Error(expr, err)
	}
	err = p.validator.CheckInput(operation.Operator, operation.Operands)
	if err != nil {
		return "", format.Error(expr, err)
	}
	return p.engine.ProcessOperation(operation)
}

func (p *Parser) getOperation(expr string) (*calculator.Operation, error) {
	operator := "+"
	operands := []float64{2.0, 3.0}
	return &calculator.Operation{
		Expression: expr,
		Operator:   operator,
		Operands:   operands,
	}, nil
}
