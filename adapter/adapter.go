package adapter

import "design_pattern/adapter/task"

// 适配器, 在不更改另外两个包的情况下, 适配这两个包
type RunnableAdapter struct {
	callable task.Callable
}

func NewRunnableAdapter(callable task.Callable) *RunnableAdapter {
	return &RunnableAdapter{callable: callable}
}

func (r *RunnableAdapter) Run() {
	r.callable.Call()
}
