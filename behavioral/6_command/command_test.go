package command_test

import (
	"command"
	"testing"
)

func TestCommand(t *testing.T) {
	saveButton := command.NewSaveButton()
	saveCommand := command.NewSaveCommand()
	saveButton.Bind(saveCommand)
	saveButton.Press()
}
