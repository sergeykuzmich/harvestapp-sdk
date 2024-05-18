package hrvst

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

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
	var notFound *httpErrors.NotFound
	ok := errors.As(err, &notFound)
	assert.True(t, ok)
}

func TestCreateTask(t *testing.T) {
	client := testClient()

	validTask := &Task{
		Name: "New Task Name",
	}

	task, err := client.CreateTask(validTask, Defaults())
	assert.Nil(t, err)

	assert.NotNil(t, task)
	assert.Equal(t, "New Task Name", task.Name)
	assert.Equal(t, 8083782, task.ID)
}

func TestCreateInvalidTask(t *testing.T) {
	client := testClient()

	invalidTask := &Task{
		DefaultHourlyRate: 120.0,
	}

	args := Arguments{}
	args["status"] = "422"

	_, err := client.CreateTask(invalidTask, args)
	assert.NotNil(t, err)

	// Since SDK uses own errors check correct type is returned
	var unprocessableEntity *httpErrors.UnprocessableEntity
	ok := errors.As(err, &unprocessableEntity)
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

	validTask := &Task{
		ID:                8083782,
		DefaultHourlyRate: 120.0,
	}

	task, err := client.UpdateTask(validTask, Defaults())
	assert.Nil(t, err)

	assert.NotNil(t, task)
	assert.Equal(t, 120.0, task.DefaultHourlyRate)
	assert.Equal(t, 8083782, task.ID)
}

func TestUpdateTaskWithInvalidInput(t *testing.T) {
	client := testClient()

	invalidTask := &Task{
		ID:   8083783,
		Name: "",
	}

	args := Arguments{}
	args["status"] = "422"

	_, err := client.UpdateTask(invalidTask, args)
	assert.NotNil(t, err)

	// Since SDK uses own errors check correct type is returned
	var unprocessableEntity *httpErrors.UnprocessableEntity
	ok := errors.As(err, &unprocessableEntity)
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
	var notFound *httpErrors.NotFound
	ok := errors.As(err, &notFound)
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

	var notFound *httpErrors.NotFound
	ok := errors.As(err, &notFound)
	assert.True(t, ok)
}

func TestGetTasksPaginated(t *testing.T) {
	client := testClient()

	var all []*Task

	tasks, next, err := client.GetTasks(Defaults())
	assert.Nil(t, err)
	assert.NotNil(t, next)
	assert.Equal(t, 3, len(tasks))

	assert.Equal(t, "Business Development", tasks[0].Name)
	assert.Equal(t, "Research", tasks[1].Name)
	assert.Equal(t, "Project Management", tasks[2].Name)

	all = append(all, tasks...)

	tasks, next, err = next()
	assert.Nil(t, err)
	assert.Nil(t, next)
	assert.Equal(t, 2, len(tasks))

	assert.Equal(t, "Programming", tasks[0].Name)
	assert.Equal(t, "Graphic Design", tasks[1].Name)

	all = append(all, tasks...)

	assert.Equal(t, 5, len(all))

	assert.Equal(t, "Business Development", all[0].Name)
	assert.Equal(t, "Research", all[1].Name)
	assert.Equal(t, "Project Management", all[2].Name)
	assert.Equal(t, "Programming", all[3].Name)
	assert.Equal(t, "Graphic Design", all[4].Name)
}
