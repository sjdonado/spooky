package test_cases

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/url"
)

type Param struct {
	Key      string   `json:"key"`
	Value    string   `json:"value"`
	Generate bool     `json:"generate"`
	Aliases  []string `json:"aliases"`
}

type ResponseValue struct {
	Key      string   `json:"key"`
	Required bool     `json:"required"`
	Match    bool     `json:"match"`
	Aliases  []string `json:"aliases"`
}

type TestCase struct {
	Name        string          `json:"name"`
	EndpointKey string          `json:"endpoint_key"`
	Method      string          `json:"method"`
	Auth        bool            `json:"auth"`
	Params      []Param         `json:"params"`
	StatusCode  int             `json:"status_code"`
	Response    []ResponseValue `json:"response"`
}

const TEMP_EMAIL = "tnfssivcxdsstil@nucleant.org"
const PASSWORD = "a6daca75e93cd05c48ba6093c0060ab7"

var TestCasesDict map[string]*TestCase

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Read() {
	file, err := ioutil.ReadFile("./input/test-cases.json")
	check(err)

	data := []TestCase{}

	err = json.Unmarshal([]byte(file), &data)
	check(err)

	TestCasesDict = make(map[string]*TestCase)
	for _, testCase := range data {
		TestCasesDict[testCase.EndpointKey] = &testCase
	}
}

func GetTestCaseByKey(key string) *TestCase {
	return TestCasesDict[key]
}

func (t TestCase) GetURLParams() url.Values {
	values := make(map[string][]string)

	for _, param := range t.Params {
		value := getParamValue(param)
		values[param.Key] = []string{value}

		if len(param.Aliases) > 0 {
			for _, alias := range param.Aliases {
				values[alias] = []string{value}
			}
		}
	}

	return values
}

func (t TestCase) GetJSONPayload() (*bytes.Buffer, error) {
	values := t.paramsToDict()
	json_data, err := json.Marshal(values)

	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(json_data), nil
}

func (t TestCase) ValidateResponse(resp map[string]string) bool {
	sent_params := t.paramsToDict()

	for _, test_resp := range t.Response {
		if test_resp.Required && resp[test_resp.Key] == "" {
			return false
		}
		if test_resp.Required && test_resp.Match {
			if resp[test_resp.Key] != sent_params[test_resp.Key] {
				return false
			}
		}
	}

	return true
}

func (t TestCase) paramsToDict() map[string]string {
	values := make(map[string]string)

	for _, param := range t.Params {
		value := getParamValue(param)

		values[param.Key] = value

		if len(param.Aliases) > 0 {
			for _, alias := range param.Aliases {
				values[alias] = value
			}
		}
	}

	return values
}

func getParamValue(param Param) string {
	if param.Generate {
		switch param.Key {
		case "email":
			return TEMP_EMAIL
		case "password":
			return PASSWORD
		}
	}
	return param.Value
}
