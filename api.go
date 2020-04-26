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

// Harvest API Client instance
type API struct {
	client      *http.Client
	apiURL      string
	AccountID   string
	AccessToken string
}

// Initialize Harvest API Client with auth credentials:
//	* Account ID
//	* Api Token.
func Client(accountID string, accessToken string) *API {
	a := API{}
	a.client = http.DefaultClient
	a.apiURL = "https://" + harvestDomain + "/" + harvestAPIVersion
	a.AccountID = accountID
	a.AccessToken = accessToken
	return &a
}

// Applies relevant User-Agent, Accept & Authorization.
func (a *API) addHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", "github.com/sergeykuzmich/harvest-sdk v"+clientVersion)
	req.Header.Set("Harvest-Account-Id", a.AccountID)
	req.Header.Set("Authorization", "Bearer "+a.AccessToken)
}

// Decode respose JSON to provided target interface.
func (a *API) decodeBody(jsonBody []byte, target interface{}) error {
	err := json.Unmarshal(jsonBody, target)
	if err != nil {
		return errors.Wrapf(err, "JSON decode failed: `%s`", string(jsonBody))
	}

	return nil
}

// Prepare & fill http.Request with URI, query, body & headers.
func (a *API) createRequest(method string, path string, queryData Arguments, postData interface{}) *http.Request {
	url := fmt.Sprintf("%s%s", a.apiURL, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, queryData.toURLValues().Encode())

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
		return httpErrors.CreateFromResponse(res)
	}

	if target != nil {
		body, _ := ioutil.ReadAll(res.Body)

		return a.decodeBody(body, target)
	}

	return nil
}

// Perform GET request to Harvest API with:
//	* path		- https://API.harvestapp.com/v2/{path}
//	* args		- as query variables
//	* target	- interface response should be placed to
func (a *API) Get(path string, args Arguments, target interface{}) error {
	req := a.createRequest("GET", path, args, nil)

	return a.doRequest(req, target)
}

// Perform DELETE request to Harvest API with:
//	* path	- https://API.harvestapp.com/v2/{path}
//	* args	- as query variables
func (a *API) Delete(path string, args Arguments) error {
	req := a.createRequest("DELETE", path, args, nil)

	return a.doRequest(req, nil)
}

// Perform POST request to Harvest API with:
//	* path		- https://API.harvestapp.com/v2/{path}
//	* args		- as query variables
//	* body		- as body
//	* target	- interface response should be placed to
func (a *API) Post(path string, args Arguments, body interface{}, target interface{}) error {
	req := a.createRequest("POST", path, args, body)

	return a.doRequest(req, target)
}

// Perform PATCH request to Harvest API with:
//	* path		- https://API.harvestapp.com/v2/{path}
//	* args		- as query variables
//	* body		- as body
//	* target	- interface response should be placed to
func (a *API) Patch(path string, args Arguments, body interface{}, target interface{}) error {
	req := a.createRequest("PATCH", path, args, body)

	return a.doRequest(req, target)
}
