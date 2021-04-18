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

type Score struct {
	Endpoint Endpoint `json:"endpoint"`
	Value    int      `json:"value"`
	Response string   `json:"response"`
}

type Project struct {
	Id        int        `json:"id"`
	Url       string     `json:"url"`
	Students  []Student  `json:"students"`
	Endpoints []Endpoint `json:"endpoints"`
	Scores    []Score    `json:"scores"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var Projects []*Project

func Read() {
	file, err := ioutil.ReadFile("./input/projects.json")
	check(err)

	err = json.Unmarshal([]byte(file), &Projects)
	check(err)
}
