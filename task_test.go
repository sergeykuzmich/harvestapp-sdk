package hrvst

import (
	"testing"

	"github.com/stretchr/testify/assert"

	http_errors "github.com/sergeykuzmich/harvestapp-sdk/http_errors"
)

func TestGetExistingTask(t *testing.T) {
	harvest := harvestTestClient()

	task, err := harvest.GetTask(8083800, Defaults())
	assert.Nil(t, err)

	assert.NotNil(t, task)
	assert.Equal(t, "Business Development", task.Name)
	assert.Equal(t, 8083800, task.ID)
}

func TestGetNonExistingTask(t *testing.T) {
	harvest := harvestTestClient()

	_, err := harvest.GetTask(404, Defaults())
	assert.NotNil(t, err)

	// Since SDK uses own errors check correct type is returned
	_, ok := err.(*http_errors.NotFound)
	assert.True(t, ok)
}

func TestCreateTask(t *testing.T) {
	harvest := harvestTestClient()

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
	harvest := harvestTestClient()

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

func TestUpdateTaskWithValidInput(t *testing.T) {
	harvest := harvestTestClient()

	valid_task := &Task{
		ID:                8083782,
		DefaultHourlyRate: 120.0,
	}

	task, err := harvest.UpdateTask(valid_task, Defaults())
	assert.Nil(t, err)

	assert.NotNil(t, task)
	assert.Equal(t, 120.0, task.DefaultHourlyRate)
	assert.Equal(t, 8083782, task.ID)
}

func TestUpdateTaskWithInvalidInput(t *testing.T) {
	harvest := harvestTestClient()

	invalid_task := &Task{
		ID:   8083783,
		Name: "",
	}

	args := Arguments{}
	args["status"] = "422"

	_, err := harvest.UpdateTask(invalid_task, args)
	assert.NotNil(t, err)

	// Since SDK uses own errors check correct type is returned
	_, ok := err.(*http_errors.UnprocessableEntity)
	assert.True(t, ok)
}

func TestUpdateNonExistingTask(t *testing.T) {
	harvest := harvestTestClient()

	task := &Task{
		ID:   404,
		Name: "Management",
	}

	_, err := harvest.UpdateTask(task, Defaults())
	assert.NotNil(t, err)

	// Since SDK uses own errors check correct type is returned
	_, ok := err.(*http_errors.NotFound)
	assert.True(t, ok)
}

func TestDeleteTask(t *testing.T) {
	harvest := harvestTestClient()

	err := harvest.DeleteTask(8083782, Defaults())
	assert.Nil(t, err)
}

func TestDeleteNonExistingTask(t *testing.T) {
	harvest := harvestTestClient()

	err := harvest.DeleteTask(404, Defaults())
	assert.NotNil(t, err)

	_, ok := err.(*http_errors.NotFound)
	assert.True(t, ok)
}
