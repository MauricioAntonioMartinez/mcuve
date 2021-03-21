package vendors

import "strings"

type Google struct {
	Domain string
}

func NewGoogle() *Google {
	return &Google{Domain: "https://www.google.com"}
}

func (g *Google) GetBaseUrl() string {
	return g.Domain
}

func (g *Google) KeywordPath(query string) string {
	return "/search?q=" + strings.Replace(query, " ", "+", -1)
}
