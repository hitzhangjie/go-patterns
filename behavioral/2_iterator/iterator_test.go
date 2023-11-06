package iterator_test

import (
	"fmt"
	"iterator"
	"testing"
)

func TestIterator(t *testing.T) {
	students := []*iterator.Student{
		&iterator.Student{Name: "zhang", Age: 100},
		&iterator.Student{Name: "wang", Age: 99},
	}
	team := &iterator.Team{
		Students: students,
	}
	it := team.Iterator()
	for it.HasMore() {
		student := it.Next()
		fmt.Println(student)
	}
}
