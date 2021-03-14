package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	configCmd = &cobra.Command{
		Use:   "conf",
		Short: "List and modifie cli configuration",
		Long:  "List and modifie cli configuration",
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println(countries)
		},
	}
)

func argsConfig() {
	configCmd.PersistentFlags().BoolVarP(&contries, "contries", "c", false, "List all the available search contries.")

}
