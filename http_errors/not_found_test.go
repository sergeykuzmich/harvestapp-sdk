package http_errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNotFoundError(t *testing.T) {
	path := "/ping"
	errorMessage := fmt.Sprintf("Not Found: %s", path)
	var body []byte


	var err HttpError
	err = createNotFound(path, body)

	err, ok := err.(*NotFound)
	assert.True(t, ok)

	assert.Equal(t, err.Path(), path)
	assert.Equal(t, err.Details(), body)
	assert.Equal(t, err.Error(), errorMessage)
}
