package hrvst

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sergeykuzmich/harvestapp-sdk/flags"
	httpErrors "github.com/sergeykuzmich/harvestapp-sdk/http_errors"
)

func TestGetExistingTask(t *testing.T) {
	client := testClient()

	task, err := client.GetTask(8083800, Defaults())
	assert.Nil(t, err)

	assert.NotNil(t, task)
	assert.Equal(t, "Business Development", task.Name)
	assert.Equal(t, 8083800, task.ID)
}

func TestGetNonExistingTask(t *testing.T) {
	client := testClient()

	_, err := client.GetTask(404, Defaults())
	assert.NotNil(t, err)

	// Since SDK uses own errors check correct type is returned
	_, ok := err.(*httpErrors.NotFound)
	assert.True(t, ok)
}

func TestCreateTask(t *testing.T) {
	client := testClient()

	valid_task := &Task{
		Name: "New Task Name",
	}

	task, err := client.CreateTask(valid_task, Defaults())
	assert.Nil(t, err)

	assert.NotNil(t, task)
	assert.Equal(t, "New Task Name", task.Name)
	assert.Equal(t, 8083782, task.ID)
}

func TestCreateInvalidTask(t *testing.T) {
	client := testClient()

	invalid_task := &Task{
		DefaultHourlyRate: 120.0,
	}

	args := Arguments{}
	args["status"] = "422"

	_, err := client.CreateTask(invalid_task, args)
	assert.NotNil(t, err)

	// Since SDK uses own errors check correct type is returned
	_, ok := err.(*httpErrors.UnprocessableEntity)
	assert.True(t, ok)

	// The way to check error details:
	//
	//	asUnprocessableEntityError, ok := err.(*httpErrors.UnprocessableEntity)
	//	assert.True(t, ok)
	//	expectedDetails := `{
	//		"message": "Name can't be blank"
	//	}`
	//	assert.Equal(t, asUnprocessableEntityError.Details(), expectedDetails)
}

func TestUpdateTaskWithValidInput(t *testing.T) {
	client := testClient()

	valid_task := &Task{
		ID:                8083782,
		DefaultHourlyRate: 120.0,
	}

	task, err := client.UpdateTask(valid_task, Defaults())
	assert.Nil(t, err)

	assert.NotNil(t, task)
	assert.Equal(t, 120.0, task.DefaultHourlyRate)
	assert.Equal(t, 8083782, task.ID)
}

func TestUpdateTaskWithInvalidInput(t *testing.T) {
	client := testClient()

	invalid_task := &Task{
		ID:   8083783,
		Name: "",
	}

	args := Arguments{}
	args["status"] = "422"

	_, err := client.UpdateTask(invalid_task, args)
	assert.NotNil(t, err)

	// Since SDK uses own errors check correct type is returned
	_, ok := err.(*httpErrors.UnprocessableEntity)
	assert.True(t, ok)
}

func TestUpdateNonExistingTask(t *testing.T) {
	client := testClient()

	task := &Task{
		ID:   404,
		Name: "Management",
	}

	_, err := client.UpdateTask(task, Defaults())
	assert.NotNil(t, err)

	// Since SDK uses own errors check correct type is returned
	_, ok := err.(*httpErrors.NotFound)
	assert.True(t, ok)
}

func TestDeleteTask(t *testing.T) {
	client := testClient()

	err := client.DeleteTask(8083782, Defaults())
	assert.Nil(t, err)
}

func TestDeleteNonExistingTask(t *testing.T) {
	client := testClient()

	err := client.DeleteTask(404, Defaults())
	assert.NotNil(t, err)

	_, ok := err.(*httpErrors.NotFound)
	assert.True(t, ok)
}

func TestGetTasksPaginated(t *testing.T) {
	client := testClient()

	var all []*Task

	tasks, next, err := client.GetTasks(Defaults())
	assert.Nil(t, err)
	assert.NotNil(t, next)
	assert.Equal(t, 3, len(tasks))

	all = append(all, tasks...)

	tasks, next, err = next()
	assert.Nil(t, err)
	assert.Nil(t, next)
	assert.Equal(t, 2, len(tasks))

	all = append(all, tasks...)

	assert.Equal(t, 5, len(all))
}

func TestGetTasksNonPaginated(t *testing.T) {
	client := testClient()

	args := Defaults()
	args[flags.GetAll] = "true"
	tasks, next, err := client.GetTasks(args)
	assert.Nil(t, err)
	assert.Nil(t, next)
	assert.Equal(t, 5, len(tasks))
}
