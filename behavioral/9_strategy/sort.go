package strategy

import "fmt"

// The Sorter interface declares a method for sorting
type Sorter interface {
	Sort(values []int)
}

// BubbleSort implements bubble sort
type BubbleSort struct{}

func (b *BubbleSort) Sort(values []int) {
	// Bubble sort implementation
	fmt.Println("Bubble Sorting...")
}

// QuickSort implements quicksort
type QuickSort struct{}

func (q *QuickSort) Sort(values []int) {
	// Quicksort implementation
	fmt.Println("Quick Sorting...")
}

func NewSorterContext(sorter Sorter) *SorterContext {
	return &SorterContext{sorter}
}

// The SorterContext holds a sorter and delegates sorting to it
type SorterContext struct {
	sorter Sorter
}

func (s *SorterContext) ExecuteSort(values []int) {
	s.sorter.Sort(values)
}
