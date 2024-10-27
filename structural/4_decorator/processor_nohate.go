package decorator

import (
	"bytes"
	"errors"
)

type nohateArticleProcessor struct {
	base Processor
}

func NewNoHateArticleProcessor(opts Options) Processor {
	return &nohateArticleProcessor{
		base: NewArticleProcessor(opts),
	}
}

func (p nohateArticleProcessor) Process(dat []byte) ([]byte, error) {
	// decorate this normal article processor, we add checking hatewords logic,
	//
	// nohateArticleProcessor and ArticleProcessor are likely to be defined
	// in different modules or packages, we cannot change ArticleProcessor's
	// implementation, so we use decorate pattern to augment it's behavior.
	if p.hasHateWords(dat) {
		return nil, errors.New("contains hate words")
	}
	return p.base.Process(dat)
}

func (p nohateArticleProcessor) hasHateWords(dat []byte) bool {
	if bytes.Contains(dat, []byte("hate")) {
		return true
	}
	if bytes.Contains(dat, []byte("kill")) {
		return true
	}
	return false
}
