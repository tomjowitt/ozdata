package mockingbird

import (
// "log"
// "strconv"
)

const (
	version = "0.0.1"
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

// func (data *SuburbData) GetSuburbByPostcode(postcode string) (sub Suburb, err error) {
// 	postcodeInt, err := strconv.ParseInt(postcode, 10, 64)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, v := range data.Suburbs {
// 		if v.Postcode == postcodeInt {
// 			for _, s := range data.States {
// 				if s.Code == v.StateCode {
// 					v.State = s
// 					return v, err
// 				}
// 			}
// 		}
// 	}
// 	log.Fatal("No suburb data")
// 	return Suburb{}, err
// }
