package main

import (
	"log"

	"github.com/sjdonado/spooky/projects"
	"github.com/sjdonado/spooky/test_cases"
)

func main() {
	projects := projects.Read()
	test_cases := test_cases.Read()

	log.Printf("%+v\n", projects)
	log.Printf("%+v\n", test_cases)
}
