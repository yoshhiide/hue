package infrastructures

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	GET  = "GET"
	PUT  = "PUT"
	POST = "POST"
)

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type HueBridgeClient struct {
	httpClient httpClient
	BridgeURL  string
	Token      string
}

func NewHueBridgeClient(url, token string) *HueBridgeClient {
	return &HueBridgeClient{
		httpClient: http.DefaultClient,
		BridgeURL:  url,
		Token:      token,
	}
}

func (cli *HueBridgeClient) PutChangeLightStatus(lightNo string, status bool) ([]byte, error) {
	params, err := json.Marshal(map[string]bool{
		"on": status,
	})
	if err != nil {
		return nil, fmt.Errorf(`params, err := json.Marshal: %w`, err)
	}
	url := fmt.Sprintf("%s/%s/%s/%s/%s/%s", cli.BridgeURL, "api", cli.Token, "lights", lightNo, "state")
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(params))
	if err != nil {
		return nil, fmt.Errorf(`res, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(params)): %w`, err)
	}

	res, err := cli.do(req)
	if err != nil {
		return nil, fmt.Errorf(`res, err := cli.do(req): %w`, err)
	}

	return res, nil
}

func (cli *HueBridgeClient) do(req *http.Request) ([]byte, error) {
	req.Header.Add("Content-Type", "application/json")
	res, err := cli.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf(`res, err := cli.httpClient.Do(req): %w`, err)
	}
	defer res.Body.Close()
	byteArray, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf(`byteArray, err := ioutil.ReadAll(res.Body): %w`, err)
	}
	return byteArray, nil
}
