package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"

	"github.com/sergeykuzmich/harvestapp-sdk/http_errors"
)

const CLIENT_VERSION = "1.0.0"
const HARVEST_DOMAIN = "api.harvestapp.com"
const HARVEST_API_VERSION = "v2"

type API struct {
	client      *http.Client
	apiUrl      string
	AccountId   string
	AccessToken string
}

func Harvest(accountId string, accessToken string) *API {
	a := API{}
	a.client = http.DefaultClient
	a.apiUrl = "https://" + HARVEST_DOMAIN + "/" + HARVEST_API_VERSION
	a.AccountId = accountId
	a.AccessToken = accessToken
	return &a
}

// Applies relevant User-Agent, Accept & Authorization
func (a *API) addHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", "github.com/sergeykuzmich/harvest-sdk v"+CLIENT_VERSION)
	req.Header.Set("Harvest-Account-Id", a.AccountId)
	req.Header.Set("Authorization", "Bearer "+a.AccessToken)
}

// Decode respose JSON to provided target interface
func (a *API) decodeBody(jsonBody []byte, target interface{}) error {
	err := json.Unmarshal(jsonBody, target)
	if err != nil {
		return errors.Wrapf(err, "JSON decode failed: %s", string(jsonBody))
	}

	return nil
}

func (a *API) createRequest(method string, path string, args Arguments, postData interface{}) *http.Request {
	url := fmt.Sprintf("%s%s", a.apiUrl, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, args.ToURLValues().Encode())

	buffer := new(bytes.Buffer)
	if postData != nil {
		json.NewEncoder(buffer).Encode(postData)
	}

	req, _ := http.NewRequest(method, urlWithParams, buffer)
	a.addHeaders(req)

	return req
}

func (a *API) doRequest(req *http.Request, target interface{}) error {
	res, err := a.client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "HTTP request failure on %s", req.URL.Path)
	}

	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return http_errors.CreateFromResponse(res)
	}

	body, _ := ioutil.ReadAll(res.Body)

	return a.decodeBody(body, target)
}

func (a *API) Get(path string, args Arguments, target interface{}) error {
	req := a.createRequest("GET", path, args, nil)

	return a.doRequest(req, target)
}

func (a *API) Post(path string, args Arguments, postData interface{}, target interface{}) error {
	req := a.createRequest("POST", path, args, postData)

	return a.doRequest(req, target)
}
