package command

import "fmt"

type Shortcut interface {
	Bind(Command)
	Press()
}

type shortcut struct {
	keys []int32
	cmd  Command
}

func NewShortcut(keys []int32) *shortcut {
	return &shortcut{keys: keys}
}

func (s *shortcut) Bind(c Command) {
	s.cmd = c
}

func (s *shortcut) Press() {
	fmt.Println("press [save] shortcut")
	s.cmd.Execute()
}
