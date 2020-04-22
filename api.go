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
	req.Header.Set("User-Agent", "github.com/sergeykuzmich/harvest-sdk v" + CLIENT_VERSION)
	req.Header.Set("Harvest-Account-Id", a.AccountId)
	req.Header.Set("Authorization", "Bearer " + a.AccessToken)
}

func (a *API) createRequest(method string, path string, args Arguments, postData interface{}) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", a.apiUrl, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, args.ToURLValues().Encode())

	buffer := new(bytes.Buffer)
	if postData != nil {
		json.NewEncoder(buffer).Encode(postData)
	}

	req, err := http.NewRequest(method, urlWithParams, buffer)
	if err != nil {
		return req, err
	}

	a.addHeaders(req)
	return req, nil
}

func (a *API) makeRequest(req *http.Request, target interface{}) error {
	res, err := a.client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "HTTP request failure on %s", req.URL.Path)
	}

	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		err := errors.New(strconv.Itoa(res.StatusCode))
		return errors.Wrapf(err, "HTTP request failure on %s", req.URL.Path)
	}

	body, err := ioutil.ReadAll(res.Body)

	return a.decodeBody(body, target)
}

func (a *API) decodeBody(jsonBody []byte, target interface{}) error {
	err := json.Unmarshal(jsonBody, target)
	if err != nil {
		return errors.Wrapf(err, "JSON decode failed: %s", string(jsonBody))
	}

	return nil
}

func (a *API) Get(path string, args Arguments, target interface{}) error {
	req, err := a.createRequest("GET", path, args, nil)
	if err !=nil {
		return err
	}

	return a.makeRequest(req, target)
}

func (a *API) Post(path string, args Arguments, postData interface{}, target interface{}) error {
	req, err := a.createRequest("POST", path, args, postData)
	if err !=nil {
		return err
	}

	return a.makeRequest(req, target)
}
