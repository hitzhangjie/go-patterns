package iterator

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s *Student) String() string {
	return fmt.Sprintf("student %s is %d-years-old", s.Name, s.Age)
}

type Team struct {
	Students []*Student
}

func (t *Team) Iterator() Iterator {
	return &sliceIterator{
		offset: 0,
		data:   t.Students,
	}
}

// Iterator iterator to traverse collection elements.
//
// use it like this:
//
//	if it.HasMore() {
//		   el := it.Next()
//		   // do something with `el`
//		   ...
//	}
type Iterator interface {
	HasMore() bool
	Next() *Student
}

type sliceIterator struct {
	offset int
	data   []*Student
}

func (si *sliceIterator) HasMore() bool {
	return si.offset < len(si.data)
}

func (si *sliceIterator) Next() *Student {
	if !si.HasMore() {
		return nil
	}
	v := si.data[si.offset]
	si.offset++
	return v
}
