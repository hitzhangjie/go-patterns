package command_test

import (
	"command"
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	// press the button to execute the command
	saveButton := command.NewSaveButton()
	saveCommand := command.NewSaveCommand()
	saveButton.Bind(saveCommand)
	saveButton.Press()

	fmt.Println()

	closeButton := command.NewCloseButton()
	closeCommand := command.NewCloseCommand()
	closeButton.Bind(closeCommand)
	closeButton.Press()

	fmt.Println()

	// press the shortcuts to execute the command
	const (
		keyDownCtrl  = 17
		keydownCharS = 83
	)
	saveShortcut := command.NewShortcut([]int32{keyDownCtrl, keydownCharS})
	saveShortcut.Bind(saveCommand)
	saveShortcut.Press()
}
