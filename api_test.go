package sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHarvestClient(t *testing.T) {
	AccountID := "ACCOUNTID"
	AccessToken := "TOKEN"

	harvest := Harvest(AccountID, AccessToken)

	assert.Equal(t, AccountID, harvest.AccountID)
	assert.Equal(t, AccessToken, harvest.AccessToken)
}

func TestInvalidJsonResponse(t *testing.T) {
	harvest := harvestTestClient()

	_, err := harvest.GetTask(8083801, Defaults())
	assert.NotNil(t, err)
}

func TestInvalidServerResponse(t *testing.T) {
	harvest := harvestTestClient()
	harvest.apiURL = mockUnstartedServerResponse().URL

	_, err := harvest.GetTask(8083801, Defaults())
	assert.NotNil(t, err)
}
