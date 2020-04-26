package sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHarvestClient(t *testing.T) {
	AccountId := "ACCOUNTID"
	AccessToken := "TOKEN"

	harvest := Harvest(AccountId, AccessToken)

	assert.Equal(t, AccountId, harvest.AccountId)
	assert.Equal(t, AccessToken, harvest.AccessToken)
}

func TestInvalidJsonResponse(t *testing.T) {
	harvest := HarvestTestClient()

	_, err := harvest.GetTask(8083801, Defaults())
	assert.NotNil(t, err)
}

func TestInvalidServerResponse(t *testing.T) {
	harvest := HarvestTestClient()
	harvest.apiUrl = mockUnstartedServerResponse().URL

	_, err := harvest.GetTask(8083801, Defaults())
	assert.NotNil(t, err)
}
