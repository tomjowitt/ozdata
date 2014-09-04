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

	states, err := ozdata.NewStates()
	if err != nil {
		fmt.Println(err)
		return
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
			return
		}
		ycoord, err := strconv.ParseFloat(record[11], 10)
		if err != nil {
			fmt.Println(err)
			return
		}

		coordinate := &ozdata.Coordinate{
			Lat:  xcoord,
			Long: ycoord,
		}

		postcodeInt, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}

		state, err := states.State(record[2])
		if err != nil {
			fmt.Println(err)
			return
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
		return
	}

	jsonResult := fmt.Sprintf(`{"suburbs": %s}`, jsonData)

	err = ioutil.WriteFile("data/suburbs.json", []byte(jsonResult), 0644)
	if err != nil {
		fmt.Println("WriteFile error:", err)
		return
	}

	fmt.Println("Data successfully imported to data/suburbs.json")
}
