package command

import "fmt"

type Command interface {
	Execute() error
}

func NewSaveCommand() Command {
	c := &saveCommand{}
	c.f = getCurrentOpenFileDOM()
	return c
}

type saveCommand struct {
	f *file
}

func (c *saveCommand) Execute() error {
	fmt.Println("get the receiver `f` which will do the task")
	return c.f.save()
}

type file struct{}

func (d *file) save() error {
	fmt.Println("file save the file")
	return nil
}

func getCurrentOpenFileDOM() *file {
	return &file{}
}
