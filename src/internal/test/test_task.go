package test

import (
	"sync"

	"github.com/samber/lo"
)

type CountingTask struct {
	mu sync.Mutex

	IsMarkedDone bool
	ValuePtr     *int64
	CounterPtr   *int64
}

func (t *CountingTask) MarkAsDone() {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.IsMarkedDone = true
}

func (t *CountingTask) Reset() {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.IsMarkedDone = false
	t.ValuePtr = nil
	t.CounterPtr = nil
}

func (t *CountingTask) SetValue(v int64) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.ValuePtr = lo.ToPtr(v)
}

func (t *CountingTask) Increment(v int64) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.CounterPtr = lo.ToPtr(lo.FromPtr(t.CounterPtr) + v)
}
