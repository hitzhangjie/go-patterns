package prototype_test

import (
	"prototype"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestPrototype(t *testing.T) {
	f := new(prototype.Factory)
	car1 := f.Create()
	spew.Dump(car1)

	car2 := car1.Clone()
	spew.Dump(car2)
}
