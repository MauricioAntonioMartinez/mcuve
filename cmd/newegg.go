package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

var (
	newEggDomain = ""
	newEggCmd = cobra.Command{
		Use:   "newegg",
		Short: "Searches products for newegg platform",
		 Run: VendorRun("newegg"),
	}
)


type NewEgg struct {
	domain string
}


func (y *NewEgg) getBaseUrl() string {
	return ""
}

func (y *NewEgg) keywordPath(query string) string {
	return strings.Replace(query," ","+",-1)
}
