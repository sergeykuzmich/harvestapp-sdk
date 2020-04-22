package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

const CLIENT_VERSION = "1.0.0"
const HARVEST_DOMAIN = "api.harvestapp.com"
const HARVEST_API_VERSION = "v2"

type API struct {
	Client      *http.Client
	ApiUrl      string
	AccountId   string
	AccessToken string
}

func Harvest(accountId string, accessToken string) *API {
	a := API{}
	a.Client = http.DefaultClient
	a.ApiUrl = "https://" + HARVEST_DOMAIN + "/" + HARVEST_API_VERSION
	a.AccountId = accountId
	a.AccessToken = accessToken
	return &a
}

// Applies relevant User-Agent, Accept & Authorization
func (a *API) _addHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", "github.com/sergeykuzmich/harvest-sdk v"+CLIENT_VERSION)
	req.Header.Set("Harvest-Account-Id", a.AccountId)
	req.Header.Set("Authorization", "Bearer "+a.AccessToken)
}

func (a *API) _makeRequest(method string, path string, args Arguments, postData interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s%s", a.ApiUrl, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, args.ToURLValues().Encode())

	buffer := new(bytes.Buffer)
	if postData != nil {
		json.NewEncoder(buffer).Encode(postData)
	}

	req, _ := http.NewRequest(method, urlWithParams, buffer)
	a._addHeaders(req)

	var body []byte

	res, err := a.Client.Do(req)
	if err != nil {
		return body, errors.Wrapf(err, "HTTP request failure on %s", url)
	}

	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		body, _ = ioutil.ReadAll(res.Body)
		err := errors.New(strconv.Itoa(res.StatusCode))
		return body, errors.Wrapf(err, "HTTP request failure on %s: %s", url, string(body))
	}

	body, err = ioutil.ReadAll(res.Body)
	return body, err
}

func (a *API) _decodeBody(jsonBody []byte, target interface{}) error {
	err := json.Unmarshal(jsonBody, target)
	if err != nil {
		return errors.Wrapf(err, "JSON decode failed: %s", string(jsonBody))
	}

	return nil
}

func (a *API) Get(path string, args Arguments, target interface{}) error {
	res, err := a._makeRequest("GET", path, args, nil)
	if err != nil {
		return err
	}
	return a._decodeBody(res, target)
}

func (a *API) Post(path string, args Arguments, postData interface{}, target interface{}) error {
	res, err := a._makeRequest("POST", path, args, postData)
	if err != nil {
		return err
	}
	return a._decodeBody(res, target)
}
