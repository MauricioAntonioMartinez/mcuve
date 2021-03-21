package vendors

import "strings"

type Twitter struct {
	Domain string
}

func NewTwitter() *Twitter {
	return &Twitter{Domain: "https://www.twitter.com"}
}

func (t *Twitter) GetBaseUrl() string {
	return t.Domain
}

func (t *Twitter) KeywordPath(query string) string {
	return "/search?q=" + strings.Replace(query, " ", "+", -1)
}
