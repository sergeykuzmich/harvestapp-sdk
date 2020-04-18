package sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

func (a *API) Get(path string, args Arguments, target interface{}) error {
	url := fmt.Sprintf("%s%s", a.ApiUrl, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, args.ToURLValues().Encode())

	req, _ := http.NewRequest("GET", urlWithParams, nil)
	a._addHeaders(req)

	res, err := a.Client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "HTTP request failure on %s", url)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		var body []byte
		body, _ = ioutil.ReadAll(res.Body)
		return errors.Errorf("HTTP request failure on %s: %s", url, string(body))
	}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(target)
	if err != nil {
		body, _ := ioutil.ReadAll(res.Body)
		return errors.Wrapf(err, "JSON decode failed on %s: %s", url, string(body))
	}

	return nil
}

// Applies relevant User-Agent, Accept & Authorization
func (a *API) _addHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "github.com/sergeykuzmich/harvest-sdk v"+CLIENT_VERSION)
	req.Header.Set("Harvest-Account-Id", a.AccountId)
	req.Header.Set("Authorization", "Bearer "+a.AccessToken)
}
