package hrvst

import (
	"github.com/sergeykuzmich/harvestapp-sdk/flags"
	"reflect"
	"strconv"
)

type paginationInfo struct {
	NextPage int `json:"next_page"`
}

type paginated func(interface{}) (paginated, error)

// getPaginated performs GET request with generalized pagination
// * args[flags.GetAll] = "true" - is used to get ALL tasks without breaking to pages
func (a *API) getPaginated(path string, args Arguments, target interface{}) (next paginated, err error) {
	targetValue := reflect.Indirect(reflect.ValueOf(target))

	if targetValue.FieldByName("Data").Kind() != reflect.Slice {
		panic("Target interface must have Data field")
	}

	next = func(i interface{}) (paginated, error) {
		req := a.createRequest("GET", path, args, nil)

		responseBody, err := a.doRequest(req)
		if err != nil {
			return nil, err
		}

		err = a.decodeBody(responseBody, i)
		if err != nil {
			return nil, err
		}

		page := &paginationInfo{}

		err = a.decodeBody(responseBody, page)
		if err != nil {
			return nil, err
		}

		var next_ paginated
		if page.NextPage != 0 {
			next_ = func(nextTarget interface{}) (next paginated, err error) {
				args["page"] = strconv.Itoa(page.NextPage)
				return a.getPaginated(path, args, nextTarget)
			}
		}

		return next_, err
	}

	if args[flags.GetAll] != "true" {
		return next(target)
	}

	data := targetValue.FieldByName("Data")

	for ok := (next != nil); ok; ok = (next != nil) {
		targetCopy := reflect.New(targetValue.Type())

		next, err = next(targetCopy.Interface())
		if err != nil {
			return nil, err
		}

		data = reflect.AppendSlice(data, reflect.Indirect(targetCopy).FieldByName("Data"))
	}

	targetValue.FieldByName("Data").Set(data)

	return nil, err
}
