package cmd

import "github.com/spf13/cobra"

var (
	vendors []string

	massiveCmd = cobra.Command{
		Use: "massive",
		Short: "Executes multiple searches at once",
		Run: func(cmd *cobra.Command,args []string){

			// if len(vendors)>0  { 
			// 	for _,v := range vendors{

			// 	}
			// }
		},
	}
)