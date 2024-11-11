package command

import "fmt"

// Command declares the behavior each button has
type Command interface {
	Execute() error
}

func NewSaveCommand() *saveCommand {
	c := &saveCommand{f: getCurrentOpenFileDOM()}
	return c
}

// saveCommand encapsulates the request to save a file
type saveCommand struct {
	f *file
}

func (c *saveCommand) Execute() error {
	fmt.Println("the receiver `f` which will do the `save` action")
	return c.f.save()
}

// closeCommand encapsulates the request to close a file
type closeCommand struct {
	f *file
}

func (c *closeCommand) Execute() error {
	fmt.Println("the receiver `f` which will do the `close` action")
	return c.f.close()
}

func NewCloseCommand() *closeCommand {
	return &closeCommand{f: getCurrentOpenFileDOM()}
}

// file is the receiver of the commands, it supports save, and close actions
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
