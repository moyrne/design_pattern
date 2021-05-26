package adapter_test

import (
	"design_pattern/adapter"
	"design_pattern/adapter/runnable"
	"design_pattern/adapter/task"
	"testing"
)

func TestRunnableAdapter_Run(t *testing.T) {
	thread := runnable.NewThread(adapter.NewRunnableAdapter(task.NewTask(100)))
	thread.Start()
}
