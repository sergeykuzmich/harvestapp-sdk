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

const clientVersion = "1.0.0"
const harvestDomain = "api.harvestapp.com"
const harvestAPIVersion = "v2"

type API struct {
	client      *http.Client
	apiURL      string
	AccountID   string
	AccessToken string
}

func Harvest(accountID string, accessToken string) *API {
	a := API{}
	a.client = http.DefaultClient
	a.apiURL = "https://" + harvestDomain + "/" + harvestAPIVersion
	a.AccountID = accountID
	a.AccessToken = accessToken
	return &a
}

// Applies relevant User-Agent, Accept & Authorization
func (a *API) addHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", "github.com/sergeykuzmich/harvest-sdk v"+clientVersion)
	req.Header.Set("Harvest-Account-Id", a.AccountID)
	req.Header.Set("Authorization", "Bearer "+a.AccessToken)
}

// Decode respose JSON to provided target interface
func (a *API) decodeBody(jsonBody []byte, target interface{}) error {
	err := json.Unmarshal(jsonBody, target)
	if err != nil {
		return errors.Wrapf(err, "JSON decode failed: `%s`", string(jsonBody))
	}

	return nil
}

func (a *API) createRequest(method string, path string, args Arguments, postData interface{}) *http.Request {
	url := fmt.Sprintf("%s%s", a.apiURL, path)
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
		return errors.Wrapf(err, "HTTP request failed: `%s`", req.URL.Path)
	}

	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return http_errors.CreateFromResponse(res)
	}

	if target != nil {
		body, _ := ioutil.ReadAll(res.Body)

		return a.decodeBody(body, target)
	}

	return nil
}

func (a *API) Get(path string, args Arguments, target interface{}) error {
	req := a.createRequest("GET", path, args, nil)

	return a.doRequest(req, target)
}

func (a *API) Delete(path string, args Arguments) error {
	req := a.createRequest("DELETE", path, args, nil)

	return a.doRequest(req, nil)
}

func (a *API) Post(path string, args Arguments, postData interface{}, target interface{}) error {
	req := a.createRequest("POST", path, args, postData)

	return a.doRequest(req, target)
}

func (a *API) Patch(path string, args Arguments, postData interface{}, target interface{}) error {
	req := a.createRequest("PATCH", path, args, postData)

	return a.doRequest(req, target)
}
