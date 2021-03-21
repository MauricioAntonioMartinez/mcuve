package vendors

import "strings"

type Udemy struct {
	Domain string
}

func NewUdemy() *Udemy {
	return &Udemy{Domain: "https://www.udemy.com"}
}

func (u *Udemy) GetBaseUrl() string {
	return u.Domain
}

func (y *Udemy) KeywordPath(query string) string {
	return "/results?search_query=" + strings.Replace(query, " ", "+", -1)
}
