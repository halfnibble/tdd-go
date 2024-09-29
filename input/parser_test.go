package input_test

import (
	"testing"

	"github.com/halfnibble/learn-tdd-go/calculator"
	"github.com/halfnibble/learn-tdd-go/input"
	"github.com/halfnibble/learn-tdd-go/mocks"
	"github.com/stretchr/testify/assert"
)

func TestProcessExpression(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		// Arrange
		expr := "2 + 3"
		operator := "+"
		operands := []float64{2.0, 3.0}
		expectedResult := "2 + 3 = 5.5"
		engine := mocks.NewOperationProcessor(t)
		validator := mocks.NewValidationHelper(t)
		parser := input.NewParser(engine, validator)
		validator.On("CheckInput", operator, operands).Return(nil).Once()
		engine.On("ProcessOperation", &calculator.Operation{
			Expression: expr,
			Operator:   operator,
			Operands:   operands,
		}).Return(expectedResult, nil).Once()
		// Act
		result, err := parser.ProcessExpression(expr)
		// Assert
		validator.AssertExpectations(t)
		engine.AssertExpectations(t)

		assert.Nil(t, err)
		assert.Equal(t, result, expectedResult)
	})
}
