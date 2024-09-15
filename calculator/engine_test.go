package calculator_test

import (
	"log"
	"os"
	"testing"

	"github.com/halfnibble/learn-tdd-go/calculator"
)

func TestMain(m *testing.M) {
	// Runs once per package
	setup()
	e := m.Run()
	teardown()
	doExit(e)
}

func setup() {
	// todo
	log.Println("Setting up...")
}

func teardown() {
	// todo
	log.Println("Tearing down...")
}

func doExit(code int) {
	// This function ensures teardown runs before exiting
	os.Exit(code)
}

func TestAdd(t *testing.T) {
	// Arrange
	e := calculator.Engine{}
	assert := func(x, y, expected float64) {
		actual := e.Add(x, y)
		if actual != expected {
			t.Errorf("Add(%.2f, %.2f) incorrect, expected %.2f, actual %.2f", x, y, expected, actual)
		}
	}
	t.Run("positive input", func(t *testing.T) {
		x, y := 2.5, 3.5
		expected := 6.0
		assert(x, y, expected)
	})
	t.Run("negative input", func(t *testing.T) {
		x, y := -2.5, -3.5
		expected := -6.0
		assert(x, y, expected)
	})
}

func BenchmarkAdd(b *testing.B) {
	e := calculator.Engine{}
	// run N number of times
	for i := 0; i < b.N; i++ {
		e.Add(2, 3)
	}
}
