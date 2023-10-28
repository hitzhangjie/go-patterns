package decorator

import (
	"bytes"
	"errors"
)

type perfectArticleProcessor struct {
	base Processor
}

func NewPerfectArticleProcessor(opts Options) Processor {
	p := perfectArticleProcessor{
		base: NewArticleProcessor(opts),
	}
	return p
}

func (p perfectArticleProcessor) Process(dat []byte) ([]byte, error) {
	if p.hasHateWords(dat) {
		return nil, errors.New("contains hate words")
	}
	return p.base.Process(dat)
}

func (p perfectArticleProcessor) hasHateWords(dat []byte) bool {
	if bytes.Contains(dat, []byte("hate")) {
		return true
	}
	if bytes.Contains(dat, []byte("kill")) {
		return true
	}
	return false
}
