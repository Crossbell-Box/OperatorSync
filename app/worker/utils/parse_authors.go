package utils

import (
	"fmt"
	"github.com/mmcdole/gofeed"
)

func ParseAuthors(authors []*gofeed.Person) []string {
	var parsedAuthors []string
	for _, gofeedAuthor := range authors {
		authorStr := gofeedAuthor.Name
		if gofeedAuthor.Email != "" {
			authorStr += fmt.Sprintf(" <%s>", gofeedAuthor.Email)
		}
		parsedAuthors = append(parsedAuthors, authorStr)
	}

	return parsedAuthors
}
