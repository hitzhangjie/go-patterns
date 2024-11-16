package visitor

import (
	"fmt"
	"math"

	"github.com/hitzhangjie/go-patterns/behavior/10_visitor/shapes"
)

type Visitor interface {
	Visit(Shape)
}

type Shape interface {
	Accept(Visitor)
}

type Square struct {
	// we need to access the square data, but cannot modify it
	shapes.Square
}

func (s *Square) Accept(v Visitor) {
	v.Visit(s)
}

type Circle struct {
	// we need to access the circle data, but cannot modify it
	shapes.Circle
}

func (s *Circle) Accept(v Visitor) {
	v.Visit(s)
}

// At the beginning, someone ask you a Feature Request: support calculating
// the area for different shapes.
type AreaVisitor struct {
}

func (v AreaVisitor) Visit(shape Shape) {
	switch v := shape.(type) {
	case *Square:
		fmt.Println("area of square:", math.Pow(v.Side, 2))
	case *Circle:
		fmt.Println("area of circle:", math.Pi*math.Pow(v.Radius, 2))
	default:
		panic("not supported")
	}
}

// Server days later, another one ask you a Feature Request: support calculating
// the perimeter for different shapes.
type PerimeterVisitor struct {
}

func (v *PerimeterVisitor) Visit(shape Shape) {
	switch v := shape.(type) {
	case *Square:
		fmt.Println("perimeter of square:", 4*v.Side)
	case *Circle:
		fmt.Println("perimeter of circle:", 2*math.Pi*v.Radius)
	default:
		panic("not supported")
	}
}

// you could add new behavior (support new Feature Requests) without changing
// the existed class structure. The element's `Accept(Visitor)`-> `Visitor.Visit(Shape)`
// let you could extend your application's ability without leaving the code
// a pile of shit.
//
// type XXXVisitor struct {
//
// }
//
// func (v *XXXVisitor) Visit(shape Shape) {
// ....
// }
