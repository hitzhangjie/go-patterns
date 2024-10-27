package composite

import (
	"errors"
	"fmt"
)

type File interface {
	Name() string
	Size() (int, error) // returns file size in bytes
	Add(f File) error
}

func NewFile(name string, sz int) File {
	return &file{name: name, size: sz}
}

type file struct {
	name string
	size int
}

func (f *file) Name() string {
	return f.name
}

func (f *file) Size() (int, error) {
	return f.size, nil
}

func (f *file) Add(file File) error {
	return fmt.Errorf("%s isn't folder, cannot add file", f.name)
}

func NewFolder(name string) File {
	return &folder{name: name, size: 4 << 10}
}

type folder struct {
	name  string
	size  int
	files []File
}

func (f *folder) Name() string {
	return f.name
}

func (f *folder) Size() (int, error) {
	totalsize := f.size
	for _, ff := range f.files {
		sz, err := ff.Size()
		if err != nil {
			return 0, err
		}
		totalsize += sz
	}
	return totalsize, nil
}

func (f *folder) Add(file File) error {
	for _, ff := range f.files {
		if ff.Name() == file.Name() {
			return errors.New("duplicate")
		}
	}
	f.files = append(f.files, file)
	return nil
}
