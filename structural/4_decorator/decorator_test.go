package decorator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArticleProcesso(t *testing.T) {
	var content = "hello world, some hate words here"

	// original processor
	opts := Options{true, true}
	p1 := NewArticleProcessor(opts)
	_, err := p1.Process([]byte(content))
	require.Nil(t, err)

	// augmented processor
	p2 := NewNoHateArticleProcessor(opts)
	_, err = p2.Process([]byte(content))
	require.NotNil(t, err)
}
