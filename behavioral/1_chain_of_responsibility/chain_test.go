package chain_of_responsibility_test

import (
	"fmt"
	chain_of_responsibility "responsibility"
	"testing"
	"time"
)

func TestChainOfResponsibility(t *testing.T) {
	article := &chain_of_responsibility.Article{
		Title:   "my dirty title",
		Content: "my dirty article content",
		Author:  "xiaoming",
		Date:    time.Now(),
	}
	fmt.Println("\nbefore processing, article info:\n", article.String())

	h1 := chain_of_responsibility.NewCheckAuthorProcessor()
	h2 := chain_of_responsibility.NewCheckDirtyWordsProcessor()
	h3 := chain_of_responsibility.NewFormatProcessor()

	h1.SetNext(h2)
	h2.SetNext(h3)
	h1.Handle(article)
	fmt.Println("\nafter processing, article info:\n", article.String())
}
