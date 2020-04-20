package sdk

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestGetExistingTask(t *testing.T) {
	harvest := HarvestTestClient()

	task, err := harvest.GetTask(8083800, Defaults())
	assert.Nil(t, err)

	assert.NotNil(t, task)
	assert.Equal(t, "Business Development", task.Name)
	assert.Equal(t, 8083800, task.ID)
}

func TestGetNonExistingTask(t *testing.T) {
	harvest := HarvestTestClient()

	_, err := harvest.GetTask(404, Defaults())
	assert.NotNil(t, err)

	originalError := errors.Unwrap(errors.Unwrap(err))
	assert.Equal(t, originalError.Error(), "404")
}

func TestCreateTask(t *testing.T) {
	harvest := HarvestTestClient()

	valid_task := Task{
		Name: "New Task Name",
	}

	task, err := harvest.CreateTask(&valid_task, Defaults())
	assert.Nil(t, err)

	assert.NotNil(t, task)
	assert.Equal(t, "New Task Name", task.Name)
	assert.Equal(t, 8083782, task.ID)
}

func TestCreateInvalidTask(t *testing.T) {
	harvest := HarvestTestClient()

	invalid_task := Task{
		DefaultHourlyRate: 120.0,
	}

	args := Arguments{}
	args["status"] = "422"

	_, err := harvest.CreateTask(&invalid_task, args)
	assert.NotNil(t, err)

	originalError := errors.Unwrap(errors.Unwrap(err))
	assert.Equal(t, originalError.Error(), "422")
}
