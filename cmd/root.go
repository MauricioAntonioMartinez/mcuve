package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	query      string
	configFile string
	amazon = Amazon{domain:"https://www.amazon.com"}
	
	rootCmd = &cobra.Command{
		Use:   "mcuve",
		Short: "Power cli to redirect to a social media.",
		Long:  `Power cli to redirect to a social media.`,

		Run: func(cmd *cobra.Command, args []string) {

			cmd.Help()
		},
	}

	countries = map[string]string{
		"Mexico":   "mx",
		"Island":   "ac",
		"Canada":   "ca",
		"Chile":    "cl",
		"Colombia": "co",
		"Brazil":   "br",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// mcuve
	rootCmd.PersistentFlags().StringVarP(&query, "query", "q", "", "The query or search to look for.")
	rootCmd.PersistentFlags().StringVar(&configFile, "config-file", "", "config file (default is $HOME/.mcuve.yaml)")

	// yt
	rootCmd.AddCommand(ytCmd)

	// amz
	rootCmd.AddCommand(amazonCmd)

	// config
	rootCmd.AddCommand(configCmd)
	argsConfig()
}
