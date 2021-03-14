package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

var (
	amazonCmd = &cobra.Command{
		Use:   "amz",
		Short: "Opens amazon to specified search",
		Long:  "Opens amazon to specified search",
		Run: func(cmd *cobra.Command, args []string) {

			if query != "" {
				search := strings.Join(strings.Split(query, " "), "+")
				openBrowser("https://www.amazon.com.mx/s/?field-keywords=" + search)
			}
		},
	}
)
