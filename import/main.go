package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/tomjowitt/ozdata/lib"
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
	var suburbs []*ozdata.Suburb

	stateData, err := ozdata.NewStateData()
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

		coordinate := &ozdata.Coordinate{
			Lat:  xcoord,
			Long: ycoord,
		}

		postcodeInt, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			fmt.Println(err)
		}

		state, err := stateData.GetStateByCode(record[2])
		if err != nil {
			fmt.Println(err)
		}

		suburb := &ozdata.Suburb{
			Postcode:   postcodeInt,
			Name:       strings.Title(strings.ToLower(record[1])),
			Coordinate: *coordinate,
			State:      state,
		}

		suburbs = append(suburbs, suburb)
	}

	jsonData, err := json.Marshal(suburbs)
	if err != nil {
		fmt.Println(err)
	}

	jsonResult := fmt.Sprintf(`{"suburbs": %s}`, jsonData)

	err = ioutil.WriteFile("data/data.json", []byte(jsonResult), 0644)
	if err != nil {
		fmt.Println("WriteFile error:", err)
	}

	fmt.Println("Data successfully imported to data/data.json")
}
