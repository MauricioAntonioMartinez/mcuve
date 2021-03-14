package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	ytFlags string
	ytCmd   = &cobra.Command{
		Use:   "yt",
		Short: "Opens youtube to specified search",
		Long:  "Opens youtube to specified search",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cmd.ValidArgs)
			fmt.Println(args)
			fmt.Println(query)
			if len(query) != 0 {
				openBrowser("https://www.youtube.com/results?search_query=" + query)
				return
			}

			if len(args) == 0 {
				fmt.Println("Needs a query")
				os.Exit(0)
			}
		},
	}
)

func configYoutube() {

}
