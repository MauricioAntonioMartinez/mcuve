package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

var (
	ytFlags string
	ytCmd   = &cobra.Command{
		Use:   "yt",
		Short: "Opens youtube to specified search",
		Long:  "Opens youtube to specified search",
		Run: VendorRun("youtube"),
	}
)

func configYoutube() {

}




type Youtube struct {domain string}

func (y *Youtube) getBaseUrl() string {
	return y.domain
}

func (y *Youtube) keywordPath(query string) string {
	return "/results?search_query="+strings.Replace(query," ","+",-1)
}
