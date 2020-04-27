package hrvst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHrvstClientArguments(t *testing.T) {
	AccountID := "ACCOUNTID"
	AccessToken := "TOKEN"

	client := Client(AccountID, AccessToken)

	assert.Equal(t, AccountID, client.AccountID)
	assert.Equal(t, AccessToken, client.AccessToken)
}

func TestInvalidJsonResponse(t *testing.T) {
	client := testClient()

	_, err := client.GetTask(8083801, Defaults())
	assert.NotNil(t, err)
}

func TestInvalidServerResponse(t *testing.T) {
	client := testClient()
	client.apiURL = mockUnstartedServerResponse().URL

	_, err := client.GetTask(8083801, Defaults())
	assert.NotNil(t, err)
}
