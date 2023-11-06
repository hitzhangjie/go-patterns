package memento_test

import (
	"memento"
	"testing"
)

func TestMemento(t *testing.T) {
	caretaker := &memento.Caretaker{}

	originator := &memento.Originator{}
	originator.SetState("Drawing 1")
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetState("Drawing 2")
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetState("Drawing 3")
	caretaker.AddMemento(originator.CreateMemento())

	// Undo and redo operations
	caretaker.Undo(originator) // Undo: Drawing 2
	caretaker.Undo(originator) // Undo: Drawing 1
	caretaker.Redo(originator) // Redo: Drawing 2
	caretaker.Redo(originator) // Redo: Drawing 3
	caretaker.Redo(originator) // Cannot redo further.
	caretaker.Undo(originator) // Undo: Drawing 2
}
