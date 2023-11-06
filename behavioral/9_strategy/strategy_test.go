package strategy_test

import (
	"strategy"
	"testing"
)

func TestStrategy(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	// Choose sorting strategy for context1
	context1 := strategy.NewSorterContext(&strategy.BubbleSort{})
	context1.ExecuteSort(values)

	// Create new context for different strategy
	context2 := strategy.NewSorterContext(&strategy.QuickSort{})
	context2.ExecuteSort(values)
}
