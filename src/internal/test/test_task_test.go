package test

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestCountingTaskReset(t *testing.T) {
	task := new(CountingTask)

	task.MarkAsDone()
	task.SetValue(100)
	task.Increment(50)

	assert.True(t, task.IsMarkedDone)
	assert.Equal(t, int64(100), lo.FromPtr(task.ValuePtr))
	assert.Equal(t, int64(50), lo.FromPtr(task.CounterPtr))

	task.Reset()

	assert.False(t, task.IsMarkedDone)
	assert.Nil(t, task.ValuePtr)
	assert.Nil(t, task.CounterPtr)
}
