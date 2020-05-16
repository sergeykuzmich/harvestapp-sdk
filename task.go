package hrvst

import (
	"fmt"
	"time"

	"github.com/sergeykuzmich/harvestapp-sdk/flags"
)

type tasksResponse struct {
	NextPage int     `json:"next_page"`
	Data     []*Task `json:"tasks"`
}

type tasksPaginated func() ([]*Task, tasksPaginated, error)

// Task is a struct to represent Harvest Task, performs:
//  * `struct` -> `JSON` convertion;
//  * `JSON` -> `struct` conversion.
type Task struct {
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	BillableByDefault bool      `json:"billable_by_default"`
	DefaultHourlyRate float64   `json:"default_hourly_rate"`
	IsDefault         bool      `json:"is_default"`
	IsActive          bool      `json:"is_active"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// getAllTasks forcely returns all Harvest Tasks existed on the Account
func (a *API) getAllTasks(args Arguments) (tasks []*Task, err error) {
	args[flags.GetAll] = "true"

	tasks, _, err = a.GetTasks(args)

	return tasks, err
}

// GetTasks returns list Harvest Tasks.
// * args[flags.GetAll] = "true" - is used to get ALL tasks without breaking to pages
func (a *API) GetTasks(args Arguments) (tasks []*Task, next tasksPaginated, err error) {
	var wrapper func(rawNext Paginated) ([]*Task, tasksPaginated, error)

	wrapper = func(rawNext Paginated) (tasks []*Task, next tasksPaginated, err error) {
		pagedResponse := &tasksResponse{}
		var nextPage Paginated

		if rawNext != nil {
			nextPage, err = rawNext(pagedResponse)
		} else {
			nextPage, err = a.GetPaginated("/tasks", args, pagedResponse)
		}

		tasks = pagedResponse.Data

		if nextPage != nil {
			next = func() ([]*Task, tasksPaginated, error) {
				return wrapper(nextPage)
			}
		}

		return
	}

	return wrapper(nil)
}

// GetTask returns Harvest Task with specified ID.
func (a *API) GetTask(taskID int, args Arguments) (task *Task, err error) {
	task = &Task{}
	path := fmt.Sprintf("/tasks/%v", taskID)
	err = a.Get(path, args, task)
	return task, err
}

// CreateTask creates Harvest Task equal *Task{} object.
func (a *API) CreateTask(t *Task, args Arguments) (task *Task, err error) {
	task = &Task{}
	err = a.Post("/tasks", args, t, task)
	return task, err
}

// UpdateTask performs Harvest Task update to match *Task{} object.
// * Task.ID is used to determine Harvest Task should be updated.
func (a *API) UpdateTask(t *Task, args Arguments) (task *Task, err error) {
	task = &Task{}
	path := fmt.Sprintf("/tasks/%v", t.ID)
	err = a.Patch(path, args, t, task)
	return task, err
}

// DeleteTask removes Harvest Task with specified ID.
func (a *API) DeleteTask(taskID int, args Arguments) (err error) {
	path := fmt.Sprintf("/tasks/%v", taskID)
	err = a.Delete(path, args)
	return err
}
