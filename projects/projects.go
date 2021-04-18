package projects

import (
	"encoding/json"
	"io/ioutil"
)

type Student struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Endpoint struct {
	Key  string `json:"key"`
	Path string `json:"path"`
}

type Project struct {
	Url       string     `json:"url"`
	Students  []Student  `json:"students"`
	Endpoints []Endpoint `json:"endpoints"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Read() []Project {
	file, err := ioutil.ReadFile("./input/projects.json")
	check(err)

	var data []Project

	err = json.Unmarshal([]byte(file), &data)
	check(err)

	return data
}
