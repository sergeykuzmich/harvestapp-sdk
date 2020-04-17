package main

import (
	"testing"
)

func TestTrueOutput(t *testing.T) {
	result := IsTrue(true)
	if result != true {
		t.Errorf("Expected true but got '%t'", result)
	}
}

func TestFalseOutput(t *testing.T) {
	result := IsTrue(false)
	if result != false {
		t.Errorf("Expected true but got '%t'", result)
	}
}
