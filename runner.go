package chrono

import "context"

type TaskRunner interface {
	Run(task Task)
}

type SimpleTaskRunner struct {
}

func NewDefaultTaskRunner() TaskRunner {
	return &SimpleTaskRunner{}
}

func (runner *SimpleTaskRunner) Run(task Task) {
	task.Run(context.Background())
}
