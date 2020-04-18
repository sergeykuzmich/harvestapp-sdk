package sdk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
)

func HarvestTestClient() *API {
	a := Harvest("ACCOUNTID", "TOKEN")
	a.ApiUrl = mockDynamicResponse().URL
	return a
}

func mockDynamicResponse() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		parts := []string{".", "mocks"}
		parts = append(parts, strings.Split(strings.TrimPrefix(r.URL.Path, "/"), "/")...)

		parts[len(parts)-1] = parts[len(parts)-1] + "-" + r.Method + ".json"
		filename := filepath.Join(parts...)

		if _, err := os.Stat(filename); os.IsNotExist(err) {
			http.Error(rw, fmt.Sprintf("%s doesn't exist", filename), http.StatusNotFound)
		}

		mockData, _ := ioutil.ReadFile(filename)
		rw.Write(mockData)
	}))
}

func mockUnstartedServerResponse() *httptest.Server {
	return httptest.NewUnstartedServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {}))
}
