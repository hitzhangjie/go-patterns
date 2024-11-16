package strategy

import (
	"sort"
)

// The Sorter interface declares a method for sorting
type Sorter interface {
	Sort(values []RankItem)
}

// InsertSort implements insert sort
type InsertSort struct{}

func (b *InsertSort) Sort(values []RankItem) {
	// Insert sort implementation , well, if len(values) < 12,
	// sort.Slice(values, lessfunc) indeed uses insert sort.
	sort.Slice(values, func(i, j int) bool {
		return values[i].Score > values[j].Score
	})
}

// QuickSort implements quicksort
type QuickSort struct{}

func (q *QuickSort) Sort(values []RankItem) {
	// Quicksort implementation
	sort.Slice(values, func(i, j int) bool {
		return values[i].Score > values[j].Score
	})
}

// NewRankList creates a new rank list
func NewRankList(cap int, sorter Sorter) *RankList {
	return &RankList{
		capacity: cap,
		sorter:   sorter,
	}
}

// RankItem is a ranked item in the rank list
type RankItem struct {
	Name  string
	Score int64
}

// RankList maintains a list of ranked items, it's the context of
// the strategy pattern, it will use the proper strategy to sort the
// items.
type RankList struct {
	capacity int
	items    []RankItem
	sorter   Sorter
}

func (r *RankList) AddItem(name string, score int64) {
	r.items = append(r.items, RankItem{name, score})

	r.sorter.Sort(r.items)

	if len(r.items) >= r.capacity {
		r.items = r.items[:r.capacity]
	}
}

func (r *RankList) Items() []RankItem {
	return r.items
}
