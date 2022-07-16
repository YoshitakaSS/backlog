package infra

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type APIClient struct {
	URL          string
	RequestData  interface{}
	ResponseData interface{}
}

func (a *APIClient) Get() error {
	resp, err := http.Get(a.URL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(body, &a.ResponseData)
}
