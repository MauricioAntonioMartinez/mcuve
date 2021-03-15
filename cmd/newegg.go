package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	newEggDomain = ""
	newEggCmd = &cobra.Command{
		Use:   "newegg",
		Short: "Searches products for newegg platform",
		 Run: VendorRun("newegg"),
	}
)


type NewEgg struct {
	domain string
}


func (y *NewEgg) getBaseUrl() string {
	return y.domain
}

func (y *NewEgg) keywordPath(query string) string {
	country := viper.Get("country")
	path := "/p/pl?d="+strings.Replace(query," ","+",-1)

	if country != "us" && country != nil { 
		path = "/global/mx-en/" + path
	}

	return path
}
