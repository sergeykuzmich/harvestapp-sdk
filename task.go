package hrvst

import (
	"fmt"
	"time"
)

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
