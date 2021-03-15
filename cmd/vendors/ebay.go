package vendors

import "strings"

type Ebay struct {
	Domain string
}

func (y *Ebay) GetBaseUrl() string {
	return y.Domain
}

func (y *Ebay) KeywordPath(query string) string {
	return "/sch/i.html?_nkw=" + strings.Replace(query, " ", "+", -1)
}
