package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	amazonDomain = "https://www.amazon.com"
	amazonCmd    = &cobra.Command{
		Use:   "amz",
		Short: "Opens amazon to specified search",
		Long:  "Opens amazon to specified search",
		Run: VendorRun("amazon"),
	}
)

type Amazon struct {
	domain string
}


func (y *Amazon) getBaseUrl() string {
		country := viper.Get("country")	
		if country != "us" && country != nil {
			y.domain = y.domain + "." + country.(string)
		}
	return y.domain
}

func (y *Amazon) keywordPath(query string) string {
	return "/s/?field-keywords="+ strings.Replace(query," ","+",-1)
}
