package hrvst

import (
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

type paginationInfo struct {
	NextPage int `json:"next_page"`
}

type paginated func(interface{}) (paginated, error)

// getPaginated performs GET request with generalized pagination
func (a *API) getPaginated(path string, args Arguments, target interface{}) (nextPage paginated, err error) {
	targetInstance := reflect.Indirect(reflect.ValueOf(target))

	if targetInstance.FieldByName("Data").Kind() != reflect.Slice {
		panic("`targetInstance` must implement `Data` field mapped to response json field.")
	}

	nextPage = func(i interface{}) (nextPage paginated, err error) {
		req := a.createRequest("GET", path, args, nil)

		responseBody, err := a.doRequest(req)
		if err != nil {
			return nil, err
		}

		page := &paginationInfo{}
		err = decodePaginatedBody(responseBody, i, page)
		if err != nil {
			return nil, err
		}

		if page.NextPage != 0 {
			nextPage = func(i interface{}) (nextPage paginated, err error) {
				args["page"] = strconv.Itoa(page.NextPage)
				return a.getPaginated(path, args, i)
			}
		}

		return nextPage, err
	}

	return nextPage(target)
}

// decodeBody reads response JSON to provided target interface & paginationInfo interface.
func decodePaginatedBody(jsonBody []byte, target interface{}, paginationInfo interface{}) (err error) {
	err = json.Unmarshal(jsonBody, target)
	if err != nil {
		return errors.Wrapf(err, "JSON decode failed: `%s`", string(jsonBody))
	}

	err = json.Unmarshal(jsonBody, paginationInfo)
	if err != nil {
		return errors.Wrapf(err, "JSON decode failed: `%s`", string(jsonBody))
	}

	return nil
}
