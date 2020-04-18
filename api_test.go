package sdk

import (
	"testing"
)

func TestHarvestClient(t *testing.T) {
	harvest := Harvest("ACCOUNTID", "TOKEN")
	if harvest.AccountId != "ACCOUNTID" {
		t.Errorf("AccountId expected to be ACCOUNTID but got '%s'", harvest.AccountId)
	}
	if harvest.AccessToken != "TOKEN" {
		t.Errorf("AccessToken expected to be TOKEN but got '%s'", harvest.AccessToken)
	}
}

func TestInvalidJsonResponse(t *testing.T) {
	harvest := HarvestTestClient()

	_, err := harvest.GetTask(8083801, Defaults())
	if err == nil {
		t.Fatal("Expected invalid JSON failure")
	}
}

func TestInvalidServerResponse(t *testing.T) {
	harvest := HarvestTestClient()
	harvest.ApiUrl = mockUnstartedServerResponse().URL

	_, err := harvest.GetTask(8083801, Defaults())
	if err == nil {
		t.Fatal("Expected HTTP request failure")
	}
}
