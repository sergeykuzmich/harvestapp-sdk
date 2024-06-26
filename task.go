package hrvst

import (
	"fmt"
	"time"
)

type tasksResponse struct {
	Data []*Task `json:"tasks"`
}

type tasksPaginated func() ([]*Task, tasksPaginated, error)

// Task is a struct to represent Harvest Task, performs:
//   - `struct` -> `JSON` conversion;
//   - `JSON` -> `struct` conversion.
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

// GetTasks returns list Harvest Tasks.
// * args[flags.GetAll] = "true" - is used to get ALL tasks without breaking to pages
func (a *API) GetTasks(args Arguments) (tasks []*Task, next tasksPaginated, err error) {
	var wrapper func(nextPage paginated) ([]*Task, tasksPaginated, error)

	wrapper = func(nextPage paginated) (tasks []*Task, next tasksPaginated, err error) {
		pagedResponse := &tasksResponse{}

		if nextPage != nil {
			nextPage, err = nextPage(pagedResponse)
		} else {
			nextPage, err = a.getPaginated("/tasks", args, pagedResponse)
		}

		tasks = pagedResponse.Data

		if nextPage != nil {
			next = func() ([]*Task, tasksPaginated, error) {
				return wrapper(nextPage)
			}
		}

		return tasks, next, err
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
