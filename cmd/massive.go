package cmd

import "github.com/spf13/cobra"

var (

	massiveCmd = &cobra.Command{
		Use: "massive",
		Short: "Executes multiple searches at once",
		Run: func(cmd *cobra.Command,args []string){

			if len(_vendors)>0 && query !=""  { 
				for _,v := range _vendors{
					v,_:=NewVendor(v)
					v.search()
				}
			}
		},
	}
)