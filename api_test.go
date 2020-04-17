package main

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
