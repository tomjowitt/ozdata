package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/tomjowitt/mockingbird/lib"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("data/import.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var suburbs []*mockingbird.Suburb

	stateData, err := mockingbird.NewStateData()
	if err != nil {
		fmt.Println(err)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			return
		}

		xcoord, err := strconv.ParseFloat(record[10], 10)
		if err != nil {
			fmt.Println(err)
		}
		ycoord, err := strconv.ParseFloat(record[11], 10)
		if err != nil {
			fmt.Println(err)
		}

		coordinate := &mockingbird.Coordinate{
			Lat:  xcoord,
			Long: ycoord,
		}

		postcodeInt, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			fmt.Println(err)
		}

		state, err := stateData.GetStateByCode(record[2])

		suburb := &mockingbird.Suburb{
			Postcode:   postcodeInt,
			Name:       strings.Title(strings.ToLower(record[1])),
			Coordinate: *coordinate,
			State:      state,
		}

		suburbs = append(suburbs, suburb)
	}

	jsonData, err := json.MarshalIndent(suburbs, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("data/data.json", jsonData, 0644)
	if err != nil {
		fmt.Printf("ioutil.WriteFile error: %+v", err)
	}

	fmt.Println("Data successfully imported to data/data.json")
}
