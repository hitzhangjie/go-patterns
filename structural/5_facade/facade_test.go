package facade_test

import (
	"facade"
	"testing"
)

func TestFacade(t *testing.T) {
	converter := facade.NewVideoConverter()
	converter.Convert("birthday.mov", "720p", "mp4")
}
