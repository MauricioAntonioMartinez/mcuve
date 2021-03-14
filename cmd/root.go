package cmd

import (
	"fmt"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	query       string
	userLicense string
	contries    = false

	rootCmd = &cobra.Command{
		Use:   "mcuve",
		Short: "Power cli to redirect to a social media.",
		Long:  `Power cli to redirect to a social media.`,

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cmd)
			fmt.Println(cmd)
			fmt.Println(contries)

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

	rootCmd.PersistentFlags().StringVarP(&query, "query", "q", "", "The query or search to look for.")
	rootCmd.AddCommand(ytCmd)
	rootCmd.AddCommand(amazonCmd)
	rootCmd.AddCommand(configCmd)
	argsConfig()
}

func initConfig() {
	if query != "" {
		viper.SetConfigFile(query)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal("Home dir not found.")
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".mcuve")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
