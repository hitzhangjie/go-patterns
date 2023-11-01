package command_test

import (
	"command"
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	saveButton := command.NewSaveButton()
	saveCommand := command.NewSaveCommand()
	saveButton.Bind(saveCommand)
	saveButton.Press()

	fmt.Println()

	closeButton := command.NewCloseButton()
	closeCommand := command.NewCloseCommand()
	closeButton.Bind(closeCommand)
	closeButton.Press()
}
