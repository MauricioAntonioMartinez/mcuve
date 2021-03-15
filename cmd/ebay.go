package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

var (
	ebayDomain = "https://www.ebay.com"
	ebayCmd    = &cobra.Command{
		Use:   "ebay",
		Short: "Opens ebay to specified search",
		Long:  "Opens ebay to specified search",
		Run: VendorRun("ebay"),
	}
)

type Ebay struct {
	domain string
}


func (y *Ebay) getBaseUrl() string {
	return y.domain
}

func (y *Ebay) keywordPath(query string) string {
	return "/sch/i.html?_nkw="+ strings.Replace(query," ","+",-1)
}
