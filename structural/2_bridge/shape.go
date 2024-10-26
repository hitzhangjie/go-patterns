package bridge

import "fmt"

type Shape interface {
	Name() string
	String() string
}

type Square struct {
	MyColor Color
	Edge    int // centimiters
}

func (s Square) Name() string {
	return "square"
}

func (s Square) String() string {
	return fmt.Sprintf("%s with edge=%dcm, rendered as %s", s.Name(), s.Edge, s.MyColor.String())
}

type Circle struct {
	MyColor Color
	Radius  int // centimeters
}

func (c Circle) Name() string {
	return "circle"
}

func (c Circle) String() string {
	return fmt.Sprintf("%s with edge=%dcm, rendered as %s", c.Name(), c.Radius, c.MyColor.String())
}
