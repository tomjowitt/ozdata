package main

import (
	"flag"
	"fmt"

	"github.com/tomjowitt/ozdata/lib"
)

func main() {

	var postcode = flag.Int64("p", 0, "A valid Australian postcode")
	flag.Parse()

	fmt.Println()
	fmt.Println("Welcome to Ozdata")
	fmt.Println()

	if *postcode == 0 {
		fmt.Println("Please enter a valid postcode to check")
		return
	}

	suburbs, err := ozdata.NewSuburbs()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Getting suburb By postcode:", *postcode)
	suburb, err := suburbs.Suburb(*postcode)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(suburb)
	fmt.Println()

	return
}
