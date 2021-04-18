package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sjdonado/spooky/projects"
	"github.com/sjdonado/spooky/test_cases"
)

func main() {
	projects.Read()
	test_cases.Read()

	ch := make(chan projects.Project, len(projects.Projects))

	for _, project := range projects.Projects {
		go grade(ch, project)
	}

	for i := 0; i < len(projects.Projects); i++ {
		log.Println(<-ch)
	}
}

func grade(ch chan projects.Project, project *projects.Project) {
	for _, endpoint := range project.Endpoints {
		test_case := test_cases.GetTestCaseByKey(endpoint.Key)
		score := projects.Score{
			Endpoint: endpoint,
		}

		var resp *http.Response
		var err error

		url := fmt.Sprintf("%s%s", project.Url, endpoint.Path)
		log.Printf("[project: %d] Sending to => %s", project.Id, url)

		if test_case.Method == "GET" {
			resp, err = http.Get(url)
		}

		if test_case.Method == "POST" {
			payload, params_err := test_case.GetJSONPayload()
			if params_err != nil {
				score.Response = params_err.Error()
				project.Scores = append(project.Scores, score)
				continue
			}
			resp, err = http.Post(url, "application/json", payload)
		}

		defer resp.Body.Close()

		if err != nil {
			score.Response = err.Error()
			project.Scores = append(project.Scores, score)
			continue
		}

		if resp.StatusCode != test_case.StatusCode {
			score.Response = fmt.Sprintf("Status code: %d, Expected: %d", resp.StatusCode, test_case.StatusCode)
			project.Scores = append(project.Scores, score)
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			score.Response = err.Error()
			project.Scores = append(project.Scores, score)
			continue
		}

		score.Response = string(body)

		res := make(map[string]string)
		json.Unmarshal(body, &res)

		if test_case.ValidateResponse(res) {
			score.Value = 1
		}

		project.Scores = append(project.Scores, score)
	}

	ch <- *project
}
