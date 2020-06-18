package hrvst

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"

	httpErrors "github.com/sergeykuzmich/harvestapp-sdk/http_errors"
)

const clientVersion = "1.0.0"
const harvestDomain = "api.harvestapp.com"
const harvestAPIVersion = "v2"

// API is main Harvest API Client instance of this SDK.
type API struct {
	client      *http.Client
	apiURL      string
	AccountID   string
	AccessToken string
}

// Client initializes Harvest API worker with auth credentials:
//	* Account ID;
//	* API Token.
func Client(accountID string, accessToken string) *API {
	a := API{}
	a.client = http.DefaultClient
	a.apiURL = "https://" + harvestDomain + "/" + harvestAPIVersion
	a.AccountID = accountID
	a.AccessToken = accessToken
	return &a
}

// addHeaders applies relevant User-Agent, Accept & Authorization.
func (a *API) addHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", "github.com/sergeykuzmich/harvest-sdk v"+clientVersion)
	req.Header.Set("Harvest-Account-Id", a.AccountID)
	req.Header.Set("Authorization", "Bearer "+a.AccessToken)
}

// decodeBody reads respose JSON to provided target interface.
func (a *API) decodeBody(jsonBody []byte, target interface{}) error {
	err := json.Unmarshal(jsonBody, target)
	if err != nil {
		return errors.Wrapf(err, "JSON decode failed: `%s`", string(jsonBody))
	}

	return nil
}

// createRequest prepares & fills http.Request with URI, query, body & headers.
func (a *API) createRequest(method string, path string, queryData Arguments, postData interface{}) *http.Request {
	url := fmt.Sprintf("%s%s", a.apiURL, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, queryData.toURLValues().Encode())

	buffer := new(bytes.Buffer)
	if postData != nil {
		err := json.NewEncoder(buffer).Encode(postData)
		if err != nil {
			panic(err)
		}
	}

	req, err := http.NewRequest(method, urlWithParams, buffer)
	if err != nil {
		panic(err)
	}

	a.addHeaders(req)

	return req
}

// doRequest performs http request & returns raw body
func (a *API) doRequest(req *http.Request) (body []byte, err error) {
	res, err := a.client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "HTTP request failed: `%s`", req.URL.Path)
	}

	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, httpErrors.CreateFromResponse(res)
	}

	body, err = ioutil.ReadAll(res.Body)

	return body, nil
}

// Get performs direct GET request to Harvest API with:
//	* path		- https://API.harvestapp.com/v2/{path};
//	* args		- as query variables;
//	* target	- interface response should be placed to;
//  ** - auth headers are included.
func (a *API) Get(path string, args Arguments, target interface{}) error {
	req := a.createRequest("GET", path, args, nil)

	body, err := a.doRequest(req)
	if err != nil {
		return err
	}

	return a.decodeBody(body, target)
}

// Delete performs direct DELETE request to Harvest API with:
//	* path	- https://API.harvestapp.com/v2/{path};
//	* args	- as query variables;
//  ** - auth headers are included.
func (a *API) Delete(path string, args Arguments) error {
	req := a.createRequest("DELETE", path, args, nil)

	_, err := a.doRequest(req)
	return err
}

// Shared PATCH, POST, PUT codebase.
// Used as a wrapper to make API request with body.
func (a *API) ppp(method string, path string, args Arguments, body interface{}, target interface{}) error {
	req := a.createRequest(method, path, args, body)

	responseBody, err := a.doRequest(req)
	if err != nil {
		return err
	}

	return a.decodeBody(responseBody, target)
}

// Post performs direct POST request to Harvest API with:
//	* path		- https://API.harvestapp.com/v2/{path};
//	* args		- as query variables;
//	* body		- as body;
//	* target	- interface response should be placed to;
//  ** - auth headers are included.
func (a *API) Post(path string, args Arguments, body interface{}, target interface{}) error {
	return a.ppp("POST", path, args, body, target)
}

// Patch performs direct PATCH request to Harvest API with:
//	* path		- https://API.harvestapp.com/v2/{path};
//	* args		- as query variables;
//	* body		- as body;
//	* target	- interface response should be placed to;
//  ** - auth headers are included.
func (a *API) Patch(path string, args Arguments, body interface{}, target interface{}) error {
	return a.ppp("PATCH", path, args, body, target)
}
