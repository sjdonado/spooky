package test_cases

import (
	"encoding/json"
	"io/ioutil"
)

type BodyValue struct {
	Key     string   `json:"key"`
	Value   string   `json:"value"`
	Aliases []string `json:"aliases"`
}

type ResponseValue struct {
	Key      string   `json:"key"`
	Required bool     `json:"required"`
	Match    bool     `json:"match"`
	Aliases  []string `json:"aliases"`
}

type TestCase struct {
	Name          string          `json:"name"`
	EndpointKey   string          `json:"endpoint_key"`
	Auth          bool            `json:"auth"`
	Body          []BodyValue     `json:"body"`
	AutoGenerated []string        `json:"auto_generated"`
	Response      []ResponseValue `json:"response"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Read() []TestCase {
	file, err := ioutil.ReadFile("./input/test-cases.json")
	check(err)

	var data []TestCase

	err = json.Unmarshal([]byte(file), &data)
	check(err)

	return data
}