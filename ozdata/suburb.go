package ozdata

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

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

type Suburbs struct {
	Suburbs []Suburb `json:"suburbs"`
}

// NewSuburbs initialises the postcode data from a JSON file
func NewSuburbs(filename string) (response Suburbs, err error) {

	datafile, err := ioutil.ReadFile(filename)
	if err != nil {
		return Suburbs{}, errors.New("Could not read file " + filename)
	}

	data := Suburbs{}
	err = json.Unmarshal([]byte(datafile), &data)
	if err != nil {
		return Suburbs{}, errors.New("Error reading JSON data")
	}

	return data, err
}

// GetSuburbByCode returns the suburb data for a specific postcode
func (data *Suburbs) GetSuburbByCode(postcode int64) (sub Suburb, err error) {

	for _, v := range data.Suburbs {
		if v.Postcode == postcode {
			return v, err
		}
	}

	return Suburb{}, errors.New("No suburb data for this postcode")
}
