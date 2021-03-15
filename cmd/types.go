package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/MauricioAntonioMartinez/mcuve/cmd/vendors"
	"github.com/spf13/cobra"
)

type Searcher interface {
	GetBaseUrl() string
	KeywordPath(query string) string
}

type Vendor struct {
	Searcher
}

func (v *Vendor) search() {
	url :=  v.GetBaseUrl()
	search := v.KeywordPath(query)
	openBrowser(url + search)
}

func NewVendor(name string) (*Vendor,error) { 
	var vd Searcher

	switch name { 
	case "youtube":		
		vd = &vendors.Youtube{Domain: "https://www.youtube.com"}
	case "amazon":
		vd =  &vendors.Amazon{Domain: "https://www.amazon.com"}
	case "newegg":
		vd = &vendors.NewEgg{Domain: "https://www.newegg.com"}
	case "ebay":
		vd = &vendors.Ebay{Domain: "https://www.ebay.com"}
	default:
		return nil,errors.New("Unknown vendor")
	}
	return &Vendor{vd},nil
}

func VendorRun(vendor string) func (cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		vendor,_:=NewVendor(vendor)

		if query !="" {
			vendor.search()
			return
		}

		if len(args) == 0 {
			fmt.Println("Needs a query")
			os.Exit(0)
		}
	}
}

