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

type SuburbMap map[int64][]Suburb

func (sm SuburbMap) UnmarshalJSON(b []byte) error {
	suburb := &[]Suburb{}
	err := json.Unmarshal(b, suburb)
	if err != nil {
		return err
	}

	for _, s := range *suburb {
		sm[s.Postcode] = append(sm[s.Postcode], s)
	}
	return nil
}

type Suburb struct {
	Name       string
	Postcode   int64
	Coordinate Coordinate
	State      State
}

type Suburbs struct {
	Suburbs SuburbMap `json:"suburbs"`
}

func LoadSuburbs() (response Suburbs, err error) {
	data := Suburbs{
		Suburbs: SuburbMap{},
	}

	err = json.Unmarshal([]byte(DATA), &data)
	if err != nil {
		return Suburbs{}, errors.New("Error reading JSON data")
	}

	return data, err
}

// NewSuburbs initialises the postcode data from a JSON file
func NewSuburbs(filename string) (response Suburbs, err error) {

	datafile, err := ioutil.ReadFile(filename)
	if err != nil {
		return Suburbs{}, errors.New("Could not read file " + filename)
	}

	data := Suburbs{
		Suburbs: SuburbMap{},
	}

	err = json.Unmarshal([]byte(datafile), &data)
	if err != nil {
		return Suburbs{}, errors.New("Error reading JSON data")
	}

	return data, err
}

// GetSuburbsByPostCode returns the suburbs data for a specific postcode
func (data *Suburbs) GetSuburbsByPostCode(postcode int64) (sub []Suburb, err error) {
	if suburbs, found := data.Suburbs[postcode]; found {
		return suburbs, nil
	}

	return nil, errors.New("No suburb data for this postcode")
}
