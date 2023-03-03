package task

import (
	"testing"

	"github.com/maddiesch/collector/internal/test"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestTaskWriter(t *testing.T) {
	task := new(test.CountingTask)

	w := &TaskWriter{Task: task}
	count, err := w.Write([]byte("hello"))
	assert.NoError(t, err)
	assert.Equal(t, 5, count)
	assert.Equal(t, int64(5), lo.FromPtr(task.CounterPtr))
}
