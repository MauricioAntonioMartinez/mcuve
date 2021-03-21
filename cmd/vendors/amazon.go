package vendors

import (
	"strings"

	"github.com/spf13/viper"
)

type Amazon struct {
	Domain string
}

func NewAmazon() *Amazon {
	return &Amazon{Domain: "https://www.amazon.com"}
}

func (y *Amazon) GetBaseUrl() string {
	country := viper.Get("country")
	if country != "us" && country != nil {
		y.Domain = y.Domain + "." + country.(string)
	}
	return y.Domain
}

func (y *Amazon) KeywordPath(query string) string {
	return "/s/?field-keywords=" + strings.Replace(query, " ", "+", -1)
}
