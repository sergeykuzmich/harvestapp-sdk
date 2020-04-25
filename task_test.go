package sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sergeykuzmich/harvestapp-sdk/http_errors"
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

	// Since SDK uses own errors check correct type is returned
	_, ok := err.(*http_errors.NotFound)
	assert.True(t, ok)
}

func TestCreateTask(t *testing.T) {
	harvest := HarvestTestClient()

	valid_task := &Task{
		Name: "New Task Name",
	}

	task, err := harvest.CreateTask(valid_task, Defaults())
	assert.Nil(t, err)

	assert.NotNil(t, task)
	assert.Equal(t, "New Task Name", task.Name)
	assert.Equal(t, 8083782, task.ID)
}

func TestCreateInvalidTask(t *testing.T) {
	harvest := HarvestTestClient()

	invalid_task := &Task{
		DefaultHourlyRate: 120.0,
	}

	args := Arguments{}
	args["status"] = "422"

	_, err := harvest.CreateTask(invalid_task, args)
	assert.NotNil(t, err)

	// Since SDK uses own errors check correct type is returned
	_, ok := err.(*http_errors.UnprocessableEntity)
	assert.True(t, ok)

	/* There is a way to check error details:

	err, ok := err.(*http_errors.UnprocessableEntity)
	assert.True(t, ok)
	expectedDetails = "{'message':'Name can't be blank'}"
	assert.Equal(t, err.Details(), expectedDetails)
	*/
}
