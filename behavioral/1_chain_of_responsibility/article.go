package chain_of_responsibility

import (
	"fmt"
	"time"
)

type Article struct {
	Title   string
	Content string
	Author  string
	Date    time.Time
}

func (a *Article) String() string {
	s := "Title: %s\nContent: %s\nAuthor: %s\nDate: %v"
	return fmt.Sprintf(s, a.Title, a.Content, a.Author, a.Date)
}
