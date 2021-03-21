package vendors

import (
	"strings"

	"github.com/spf13/viper"
)

type NewEgg struct {
	Domain string
}

func NewNewEgg() *NewEgg {
	return &NewEgg{Domain: "https://www.newegg.com"}
}

func (y *NewEgg) GetBaseUrl() string {
	return y.Domain
}

func (y *NewEgg) KeywordPath(query string) string {
	country := viper.Get("country")
	path := "/p/pl?d=" + strings.Replace(query, " ", "+", -1)

	if country != "us" && country != nil {
		path = "/global/mx-en/" + path
	}

	return path
}
