package visitor

import (
	"testing"
)

func TestVisitor(t *testing.T) {
	circle := &Circle{radius: 1.0}
	square := &Square{side: 2.0}

	areaVisitor := &AreaVisitor{}
	circle.Accept(areaVisitor)
	square.Accept(areaVisitor)

	perimeterVisitor := &PerimeterVisitor{}
	circle.Accept(perimeterVisitor)
	square.Accept(perimeterVisitor)
}
