package packages

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Package struct {
	License  string   `json: license`
	Name     string   `json: name`
	Summary  string   `json: summary`
	Versions []string `json: versions`
}

const fetchURL = "https://package.elm-lang.org/search.json"

func FetchPackageDatas() []*Package {
	result := []*Package{}
	resq, err := http.Get(fetchURL)
	if err != nil {
		return result
	}

	defer resq.Body.Close()

	byteArray, err := ioutil.ReadAll(resq.Body)
	if err != nil {
		return result
	}

	err = json.Unmarshal(byteArray, &result)
	if err != nil {
		return result
	}

	return result
}
