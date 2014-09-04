package ozdata

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strconv"
)

type States struct {
	States []State `json:"states"`
}

type Country struct {
	Name string
	Code string
}

type State struct {
	Name          string
	Code          string
	Capital       string
	Country       Country
	PostcodeRange []PostcodeRange `json:"PostcodeRange"`
}

type PostcodeRange struct {
	Low  int64
	High int64
}

func NewStates() (response States, err error) {

	filename := "data/states.json"

	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return States{}, errors.New("Could not find file data/states.json")
		}
	}

	datefile, err := ioutil.ReadFile(filename)
	if err != nil {
		return States{}, errors.New("Could not read file data/states.json")
	}

	data := States{}
	err = json.Unmarshal([]byte(datefile), &data)
	if err != nil {
		return States{}, errors.New("Invalid JSON in data/states.json")
	}

	return data, err
}

func (data *States) GetStates() []State {
	return data.States
}

func (data *States) GetStateByPostCode(postcode string) (state State, err error) {
	postcodeInt, err := strconv.ParseInt(postcode, 10, 64)
	if err != nil {
		return State{}, errors.New("Could not convert postcode to an integer")
	}
	for _, v := range data.States {
		for _, r := range v.PostcodeRange {
			if postcodeInt >= r.Low && postcodeInt <= r.High {
				return v, err
			}
		}
	}
	return State{}, errors.New("Invalid state code")
}

func (data *States) State(code string) (state State, err error) {
	for _, v := range data.States {
		if code == v.Code {
			return v, err
		}
	}
	return State{}, errors.New("Could not find state")
}
