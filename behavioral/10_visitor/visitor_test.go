package visitor

import (
	"fmt"
	"testing"

	"github.com/hitzhangjie/go-patterns/behavior/10_visitor/shapes"
)

func TestVisitor(t *testing.T) {
	fmt.Println("initialize the shapes")

	circle := &Circle{shapes.Circle{Radius: 1.0}}
	square := &Square{shapes.Square{Side: 2.0}}
	fmt.Println()

	fmt.Println("calculate the area of shapes:")
	areaVisitor := &AreaVisitor{}
	circle.Accept(areaVisitor)
	square.Accept(areaVisitor)
	fmt.Println()

	fmt.Println("calculate the perimeter of shapes:")
	perimeterVisitor := &PerimeterVisitor{}
	circle.Accept(perimeterVisitor)
	square.Accept(perimeterVisitor)
	fmt.Println()
}
