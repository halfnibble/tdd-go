package input

import (
	"github.com/halfnibble/learn-tdd-go/calculator"
)

type Parser struct {
	engine    *calculator.Engine
	validator *Validator
}

func (p *Parser) ProcessExpression(expression string) (*string, error) {
	return nil, nil
}
