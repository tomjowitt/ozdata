package ozdata

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Suburbs struct {
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

func NewSuburbs() (response Suburbs, err error) {

	filename := "data/data.json"

	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return Suburbs{}, errors.New("Could not find file data/data.json")
		}
	}

	datafile, err := ioutil.ReadFile(filename)
	if err != nil {
		return Suburbs{}, errors.New("Could not read file data/data.json")
	}

	data := Suburbs{}
	err = json.Unmarshal([]byte(datafile), &data)
	if err != nil {
		return Suburbs{}, errors.New("Error reading JSON data")
	}

	return data, err
}

func (data *Suburbs) Suburb(postcode int64) (sub Suburb, err error) {

	for _, v := range data.Suburbs {
		if v.Postcode == postcode {
			return v, err
		}
	}

	return Suburb{}, errors.New("No suburb data")
}
