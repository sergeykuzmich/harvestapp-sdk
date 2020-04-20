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

	}
}
