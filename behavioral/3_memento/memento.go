// Package memento provides an example of how the Memento pattern can be
// implemented in Go for a drawing tool that supports undo and redo
// functionality:
//
// This demo is really simple. In real world, it could be really complex.
// For example, drawing a line operation may consider:
// - straight line or handwritten?
// - the line color?
// - the line thickness?
// - the line length?
// - the line position? starting point and ending point.
// And a real drawing tool may support lines, circles, squares, etc.
//
// So the `originator.state` and `memento.state` may be complex and different.
// And, then the `caretaker“ will have many originators, which may create
// different mementos. Each kind of memento records the state for different
// kinds of shape.
//
// Maybe a real editor should do like this.
package memento

import "fmt"

// Memento represents the saved state of the drawing tool
type Memento struct {
	state string
}

// Originator represents the drawing tool
type Originator struct {
	state string
}

func (o *Originator) SetState(state string) {
	fmt.Println(state)
	o.state = state
}

func (o *Originator) GetState() string {
	return o.state
}

func (o *Originator) CreateMemento() *Memento {
	return &Memento{state: o.state}
}

func (o *Originator) RestoreMemento(m *Memento) {
	o.state = m.state
}

// Caretaker manages the Memento objects and provides undo and redo functionality
type Caretaker struct {
	mementos []*Memento
	current  int
}

func (c *Caretaker) AddMemento(m *Memento) {
	c.mementos = append(c.mementos, m)
	c.current = len(c.mementos) - 1
}

func (c *Caretaker) Undo(originator *Originator) {
	if c.current > 0 {
		c.current--
		originator.RestoreMemento(c.mementos[c.current])
		fmt.Println("Undo: ", originator.GetState())
	} else {
		fmt.Println("Cannot undo further.")
	}
}

func (c *Caretaker) Redo(originator *Originator) {
	if c.current < len(c.mementos)-1 {
		c.current++
		originator.RestoreMemento(c.mementos[c.current])
		fmt.Println("Redo: ", originator.GetState())
	} else {
		fmt.Println("Cannot redo further.")
	}
}
