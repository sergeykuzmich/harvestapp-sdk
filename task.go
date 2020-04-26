package hrvst

import (
	"fmt"
	"time"
)

//
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

// Get Task with specified ID
func (a *API) GetTask(taskID int, args Arguments) (task *Task, err error) {
	task = &Task{}
	path := fmt.Sprintf("/tasks/%v", taskID)
	err = a.Get(path, args, task)
	return task, err
}

// Create Task equal *Task{} object
func (a *API) CreateTask(t *Task, args Arguments) (task *Task, err error) {
	task = &Task{}
	err = a.Post("/tasks", args, t, task)
	return task, err
}

// Update Task to match *Task{} object
// > Task.ID is used to determine Harvest Task to update
func (a *API) UpdateTask(t *Task, args Arguments) (task *Task, err error) {
	task = &Task{}
	path := fmt.Sprintf("/tasks/%v", t.ID)
	err = a.Patch(path, args, t, task)
	return task, err
}

// Delete Task with specified ID
func (a *API) DeleteTask(taskID int, args Arguments) (err error) {
	path := fmt.Sprintf("/tasks/%v", taskID)
	err = a.Delete(path, args)
	return err
}
