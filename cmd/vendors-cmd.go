package cmd

import "github.com/spf13/cobra"


var (
	amazonCmd    = &cobra.Command{
		Use:   "amz",
		Short: "Opens amazon to specified search",
		Long:  "Opens amazon to specified search",
		Run: VendorRun("amazon"),
	}

	ytCmd   = &cobra.Command{
		Use:   "yt",
		Short: "Opens youtube to specified search",
		Long:  "Opens youtube to specified search",
		Run: VendorRun("youtube"),
	}

	ebayCmd    = &cobra.Command{
		Use:   "ebay",
		Short: "Opens ebay to specified search",
		Long:  "Opens ebay to specified search",
		Run: VendorRun("ebay"),
	}


	newEggCmd = &cobra.Command{
		Use:   "newegg",
		Short: "Searches products for newegg platform",
		 Run: VendorRun("newegg"),
	}
)
