package vendors

import "strings"

type FaceBook struct {
	Domain string
}

func NewFaceBook() *FaceBook {
	return &FaceBook{Domain: "https://www.facebook.com"}
}

func (f *FaceBook) GetBaseUrl() string {
	return f.Domain
}

func (f *FaceBook) KeywordPath(query string) string {
	return "/search/top?q=" + strings.Replace(query, " ", "+", -1)
}
