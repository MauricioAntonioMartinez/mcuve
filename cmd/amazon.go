package cmd

import (
	"fmt"
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
	if query != "" {
		country := viper.Get("country")
		fmt.Println(country)

		if country != "us" && country != nil {
			amazonDomain = amazonDomain + "." + country.(string)
		}
		search := strings.Join(strings.Split(query, " "), "+")
		openBrowser(fmt.Sprint(amazonDomain+"/s/?field-keywords=", search))
	}
	return ""
}

func (y *Amazon) keywordPath(query string) string {
	return strings.Replace(query," ","+",-1)
}
