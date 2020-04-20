package sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHarvestClient(t *testing.T) {
	AccountId := "ACCOUNTID"
	Token := "TOKEN"

	harvest := Harvest(AccountId, Token)

	assert.Equal(t, AccountId, harvest.AccountId)
	assert.Equal(t, Token, harvest.AccessToken)
}

func TestInvalidJsonResponse(t *testing.T) {
	harvest := HarvestTestClient()

	_, err := harvest.GetTask(8083801, Defaults())
	assert.NotNil(t, err)
}

func TestInvalidServerResponse(t *testing.T) {
	harvest := HarvestTestClient()
	harvest.ApiUrl = mockUnstartedServerResponse().URL

	_, err := harvest.GetTask(8083801, Defaults())
	assert.NotNil(t, err)
}
