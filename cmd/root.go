package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func getVendors() []string {
	return []string{"amazon", "newegg", "ebay"}
}

func getSocialNetworks() []string {
	return []string{"youtube", "twitter", "udemy", "facebook"}
}

var (
	query      string
	configFile string

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

	// newegg
	rootCmd.AddCommand(newEggCmd)

	// ebay
	rootCmd.AddCommand(ebayCmd)

	// all search
	allFlags()
	rootCmd.AddCommand(allCmd)

	rootCmd.AddCommand(youtubeCmd)

	rootCmd.AddCommand(googleCmd)

	rootCmd.AddCommand(udemyCmd)

	// config
	rootCmd.AddCommand(configCmd)
	argsConfig()
}
