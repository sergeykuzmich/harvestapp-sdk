package hrvst

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func testClient() *API {
	a := Client("ACCOUNT-ID", "TOKEN")
	a.apiURL = mockDynamicResponse().URL
	return a
}

func mockDynamicResponse() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		parts := []string{".", "_mocks"}
		parts = append(parts, strings.Split(strings.TrimPrefix(r.URL.Path, "/"), "/")...)

		parts[len(parts)-1] = parts[len(parts)-1] + "-" + r.Method

		customStatus := strings.Join(r.URL.Query()["status"], "")
		if customStatus != "" {
			parts[len(parts)-1] = parts[len(parts)-1] + "-" + customStatus

			responseStatus, err := strconv.Atoi(customStatus)
			if err != nil {
				panic(err)
			}

			rw.WriteHeader(responseStatus)
		}

		pagination := strings.Join(r.URL.Query()["page"], "")
		if pagination != "" {
			parts[len(parts)-1] = parts[len(parts)-1] + "-P" + pagination
		}

		parts[len(parts)-1] = parts[len(parts)-1] + ".json"
		filename := filepath.Join(parts...)

		if _, err := os.Stat(filename); os.IsNotExist(err) {
			http.Error(rw, fmt.Sprintf("%s doesn't exist", filename), http.StatusNotFound)
			return
		}

		mockData, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}

		rw.Write(mockData)
	}))
}

func mockUnstartedServerResponse() *httptest.Server {
	return httptest.NewUnstartedServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {}))
}
