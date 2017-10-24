package main

import (
	"flag"
	"fmt"

	"github.com/tomjowitt/ozdata/ozdata"
)

func main() {

	var postcode = flag.Int64("p", 0, "A valid 4-digit Australian postcode")
	flag.Parse()

	if *postcode == 0 {
		fmt.Println("Please enter a valid Australian postcode to check")
		return
	}

	filename := "data/suburbs.json"
	suburbs, err := ozdata.NewSuburbs(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Getting suburb By postcode:", *postcode)
	suburb, err := suburbs.GetSuburbsByPostCode(*postcode)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fmt.Sprintf("%v", suburb))
	fmt.Println()

	return
}
