package main

import (
	"fmt"
	"github.com/tomjowitt/ozdata/lib"
)

func main() {

	fmt.Println("")
	fmt.Println("Welcome to Ozdata")
	fmt.Println("")

	data, err := ozdata.NewStateData()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Raw Data:")
	fmt.Println(data)
	fmt.Println("")
	fmt.Println("Countries:")
	fmt.Println(data.GetCountries())
	fmt.Println("")
	fmt.Println("States:")
	fmt.Println(data.GetStates())
	fmt.Println("")

	var postcode string = "3200"
	fmt.Println("State By Postcode:", postcode)
	stateByPostcode, err := data.GetStateByPostCode(postcode)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stateByPostcode)
	fmt.Println("")

	// postcode = "20489"
	// fmt.Println("Suburb By Postcode:", postcode)
	// subByPostcode, err := data.GetSuburbByPostcode(postcode)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(subByPostcode)
	// fmt.Println("")

	return

}
