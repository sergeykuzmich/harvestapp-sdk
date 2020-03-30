package main

import (
	"testing"
)

func TestTrueOutput(t *testing.T) {
	result := IsTrue(true)
	if result != true {
		t.Error("Result should be true.")
	}
}

func TestFalseOutput(t *testing.T) {
	result := IsTrue(false)
	if result != false {
		t.Errorf("Expected result is false but ")
	}
}
