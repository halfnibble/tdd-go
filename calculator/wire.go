//go:build wireinject

package calculator

import "github.com/google/wire"

var Set = wire.NewSet(NewEngine, wire.Bind(new(Adder), new(*Engine)), NewCalculator)

func InitCalc() *Calculator {
	wire.Build(Set)
	return nil
}
