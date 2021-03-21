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
	url := v.GetBaseUrl()
	search := v.KeywordPath(query)
	openBrowser(url + search)
}

func NewVendor(name string) (*Vendor, error) {
	var vd Searcher

	switch name {
	case "youtube":
		vd = vendors.NewYoutube()
	case "amazon":
		vd = vendors.NewAmazon()
	case "newegg":
		vd = vendors.NewNewEgg()
	case "ebay":
		vd = vendors.NewEbay()
	case "twitter":
		vd = vendors.NewTwitter()
	case "udemy":
		vd = vendors.NewUdemy()
	case "facebook":
		vd = vendors.NewFaceBook()
	case "google":
		vd = vendors.NewGoogle()
	default:
		return nil, errors.New("Unknown vendor")
	}
	return &Vendor{vd}, nil
}

func VendorRun(vendor string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		vendor, _ := NewVendor(vendor)

		if query != "" {
			vendor.search()
			return
		}

		if len(args) == 0 {
			fmt.Println("Needs a query")
			os.Exit(0)
		}
	}
}
