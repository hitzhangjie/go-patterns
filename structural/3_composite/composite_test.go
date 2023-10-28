package composite_test

import (
	"composite"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestComposite(t *testing.T) {
	fd1 := composite.NewFolder("folder1")
	fd1.Add(composite.NewFile("f1", 1<<10))
	fd1.Add(composite.NewFile("f2", 2<<10))

	sz, err := fd1.Size()
	require.Nil(t, err)
	require.Equal(t, sz, 7<<10) // folder(4k) + f1(1K) + f2(2k) = 7K

	fd2 := composite.NewFolder("folder2")
	fd2.Add(composite.NewFile("f3", 1<<10))
	fd2.Add(composite.NewFile("f4", 2<<10)) // 7k

	fd1.Add(fd2)
	sz, err = fd1.Size()
	require.Nil(t, err)
	require.Equal(t, sz, 14<<10) // fd1+fd2 = 14k
}
