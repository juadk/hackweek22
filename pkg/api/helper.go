package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client struct
type Client struct {
	httpClient *http.Client
}

// SimpleSinglePrice https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd
type SimpleSinglePrice struct {
	ID          string
	Currency    string
	MarketPrice float32
}

// NewClient create new client object
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{httpClient: httpClient}
}

// doReq HTTP client
func doReq(req *http.Request, client *http.Client) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

// MakeReq HTTP request helper
func (c *Client) MakeReq(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	resp, err := doReq(req, c.httpClient)
	if err != nil {
		return nil, err
	}
	return resp, err
}
