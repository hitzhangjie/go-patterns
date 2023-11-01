package command

import "fmt"

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
	fmt.Println("delegate btn press event to command")
	b.command.Execute()
}

func NewSaveButton() Button {
	return &saveButton{}
}

type saveButton struct {
	baseButton
}
