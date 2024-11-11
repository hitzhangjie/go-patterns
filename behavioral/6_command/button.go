package command

import "fmt"

// Button defines the button behavior, when we press it it will generate
// a command that represents the action that we want to do.
type Button interface {
	Bind(Command)
	Press()
}

type baseButton struct {
	command Command
}

func (b *baseButton) Bind(cmd Command) {
	b.command = cmd
}

func (b *baseButton) Press() {
	b.command.Execute()
}

// NewSaveButton create he save button
func NewSaveButton() Button {
	b := &saveButton{}
	b.Text = "Save"
	return b
}

type saveButton struct {
	baseButton
	Text string
}

func (b *saveButton) Press() {
	fmt.Println("press [save] button")
	b.baseButton.Press()
}

// NewCloseButton create the close button
func NewCloseButton() Button {
	b := &closeButton{}
	b.Text = "Close"
	return b
}

type closeButton struct {
	baseButton
	Text string
}

func (b *closeButton) Press() {
	fmt.Println("press [close] button")
	b.baseButton.Press()
}
