package processor

import (
	"fmt"

	gzip "compress/gzip"
	json "encoding/json"
	ioutil "io/ioutil"
	http "net/http"
)

func Fetch(url string) (*http.Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Failed to fetch %s: %s", url, response.Status)
	}
  return response, nil
}

func FetchAndUnmarshal(url string, target interface{}) error {
  response, err := http.Get(url)
  if err != nil {
    return err
  }
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, target)
}

func FetchAndUnmarshalWithGunzip(url string, target interface{}) error {
  response, err := http.Get(url)
	reader, err := gzip.NewReader(response.Body)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, target)
}
