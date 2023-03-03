package task

type Task interface {
	Reset()

	SetValue(int64)

	MarkAsDone()

	Increment(int64)
}

type NullTask struct {
}

func (*NullTask) Reset() {}

func (*NullTask) SetValue(int64) {}

func (*NullTask) MarkAsDone() {}

func (*NullTask) Increment(int64) {}

var _ Task = (*NullTask)(nil)

type TaskWriter struct {
	Task
}

func (w *TaskWriter) Write(b []byte) (int, error) {
	w.Task.Increment(int64(len(b)))

	return len(b), nil
}
