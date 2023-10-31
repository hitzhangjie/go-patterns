package chain_of_responsibility

import (
	"fmt"
	"strings"
	"time"
)

type Handler interface {
	Handle(*Article) error
	SetNext(Handler)
}

type baseHandler struct {
	next Handler
}

func (h *baseHandler) Handle(article *Article) error {
	// preprocess
	// ...

	// call next handler
	if h.next == nil {
		return nil
	}

	if err := h.next.Handle(article); err != nil {
		return err
	}

	// postprocess
	// ...

	return nil
}

func (h *baseHandler) SetNext(hh Handler) {
	h.next = hh
}

func NewCheckAuthorProcessor() Handler {
	return &checkAuthorProcessor{}
}

type checkAuthorProcessor struct {
	baseHandler
}

func (p *checkAuthorProcessor) Handle(article *Article) error {
	// preprocess
	if article.Author == "xiaoming" {
		article.Author = "xiao****"
	}

	// bypass article to next Handler, you can also call `p.baseHandler.next.Handle(article)`
	return p.baseHandler.Handle(article)
}

func NewCheckDirtyWordsProcessor() Handler {
	return &checkDirtyWordsProcessor{}
}

type checkDirtyWordsProcessor struct {
	baseHandler
}

func (p *checkDirtyWordsProcessor) Handle(article *Article) error {
	article.Title = strings.ReplaceAll(article.Title, "dirty", "******")
	article.Content = strings.ReplaceAll(article.Content, "dirty", "******")
	return p.baseHandler.Handle(article)
}

func NewFormatProcessor() Handler {
	return &formatProcessor{}
}

type formatProcessor struct {
	baseHandler
}

func (p *formatProcessor) Handle(article *Article) error {
	article.Author = fmt.Sprintf("<h3>%s</h3>", article.Author)
	article.Content = fmt.Sprintf("<body>%s</body>", article.Content)
	article.Title = fmt.Sprintf("<h1>%s</h1", article.Title)
	article.Date = article.Date.Truncate(time.Hour)
	return p.baseHandler.Handle(article)
}
