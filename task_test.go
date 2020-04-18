package sdk

import (
	"testing"
)

func TestGetExistingTask(t *testing.T) {
	harvest := HarvestTestClient()

	task, err := harvest.GetTask(8083800, Defaults())
	if err != nil {
		t.Fatal(err)
	}

	if task == nil {
		t.Fatal("GetTask returned nil instead of task")
	}

	if task.Name != "Business Development" || task.ID != 8083800 {
		t.Errorf("Incorrect was returned")
	}
}

func TestGetNonExistingTask(t *testing.T) {
	harvest := HarvestTestClient()

	_, err := harvest.GetTask(404, Defaults())
	if err == nil {
		t.Fatal("Expected Not Found failure")
	}
}
