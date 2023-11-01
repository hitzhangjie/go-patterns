package command

import "fmt"

type Command interface {
	Execute() error
}

func NewSaveCommand() Command {
	c := &saveCommand{f: getCurrentOpenFileDOM()}
	return c
}

type saveCommand struct {
	f *file
}

func (c *saveCommand) Execute() error {
	fmt.Println("the receiver `f` which will do the task")
	return c.f.save()
}

type file struct{}

func (d *file) save() error {
	fmt.Println("receiver save the file")
	return nil
}

func (d *file) close() error {
	fmt.Println("receiver close the file")
	return nil
}

func getCurrentOpenFileDOM() *file {
	return &file{}
}

type closeCommand struct {
	f *file
}

func (c *closeCommand) Execute() error {
	fmt.Println("the receiver `f` which will do the task")
	return c.f.close()
}

func NewCloseCommand() Command {
	return &closeCommand{f: getCurrentOpenFileDOM()}
}
