package cmd

import "github.com/spf13/cobra"

var (
	typeQuery string
	allCmd    = &cobra.Command{
		Use:   "all",
		Short: "Executes multiple searches at once",
		Run: func(cmd *cobra.Command, args []string) {
			var pages []string
			if typeQuery == "vendors" {
				pages = getVendors()
			}

			if typeQuery == "social-media" {
				pages = getSocialNetworks()
			}

			if typeQuery == "" {
				cmd.Help()
				return
			}

			if len(pages) > 0 && query != "" {
				for _, v := range pages {
					v, _ := NewVendor(v)
					v.search()
				}
			}
		},
	}
)

func allFlags() {
	allCmd.PersistentFlags().StringVarP(&typeQuery, "type", "t", "vendors", "Specify which target(vendors,social-media)")
}
