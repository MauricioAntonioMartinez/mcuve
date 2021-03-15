package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	showCountries = false
	configCmd     = &cobra.Command{
		Use:   "config",
		Short: "List and modifies cli configuration",
		Run: func(cmd *cobra.Command, args []string) {

			if showCountries && len(args) == 0 {
				for ctr, key := range countries {
					fmt.Printf("[%s] %s\n", key, ctr)
				}
				return
			}

			cmd.Help()

		},
	}
)

func argsConfig() {
	configCmd.PersistentFlags().BoolVarP(&showCountries, "countries", "c", false, "List all the available search contries.")

}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal("Home dir not found.")
		}
		viper.SetConfigName(".mcuve")
		viper.AddConfigPath(home)
		viper.SetDefault("country", "us")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		createConfigFile()
	}
}

func createConfigFile() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal("Home dir not found.")
	}

	rootConfigFile := path.Join(home, ".mcuve.yaml")

	if configFile != "" {
		rootConfigFile = configFile
	}

	err = ioutil.WriteFile(rootConfigFile, make([]byte, 0), 0644)

	if err != nil {
		log.Fatalf("Cannot create config file: %s", rootConfigFile)
	}

}
