package ozdata

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type SuburbData struct {
	Suburbs []Suburb `json:"suburbs"`
}

type Coordinate struct {
	Lat  float64
	Long float64
}

type Suburb struct {
	Name       string
	Postcode   int64
	Coordinate Coordinate
	State      State
}

func NewSuburbData() (response SuburbData, err error) {

	filename := "data/data.json"

	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			log.Fatal("Could not find file data/data.json:", err)
		}
	}

	datafile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Could not read file data/data.json:", err)
	}

	data := SuburbData{}
	err = json.Unmarshal([]byte(datafile), &data)
	if err != nil {
		log.Fatal("Error reading JSON data:", err)
	}

	return data, err
}

func (data *SuburbData) GetSuburbByPostcode(postcode int64) (sub Suburb, err error) {
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range data.Suburbs {
		if v.Postcode == postcode {
			return v, err
		}
	}
	log.Fatal("No suburb data")
	return Suburb{}, err
}
