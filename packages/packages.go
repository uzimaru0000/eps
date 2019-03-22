package packages

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Package struct {
	License  string   `json: license`
	Name     string   `json: name`
	Summary  string   `json: summary`
	Versions []string `json: versions`
}

const fetchURL = "https://package.elm-lang.org/search.json"

func PackagesFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func FetchPackagesFile() ([]byte, error) {
	resq, err := http.Get(fetchURL)
	if err != nil {
		return []byte{}, err
	}

	defer resq.Body.Close()

	byteArray, err := ioutil.ReadAll(resq.Body)
	if err != nil {
		return []byte{}, err
	}

	return byteArray, nil
}

func ReadPackagesFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return []byte{}, err
	}

	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return []byte{}, err
	}

	buf := make([]byte, info.Size())
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			break
		}
	}

	return buf, nil
}

func SavePackagesFile(path string, data []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func ConverteJSON(jsonData []byte) ([]*Package, error) {
	packages := []*Package{}

	err := json.Unmarshal(jsonData, &packages)
	return packages, err
}

func CacheCheck(updateDate time.Time, date time.Time) bool {
	oneWeekSecond := time.Hour * 24 * 7
	return date.Sub(updateDate) >= oneWeekSecond
}
