package task

import "fmt"

type (
	Task struct {
		num int
	}
	Callable interface {
		Call() int
	}
)

func NewTask(num int) *Task {
	return &Task{
		num: num,
	}
}

func (t *Task) Call() (r int) {
	for i := 1; i <= t.num; i++ {
		r += i
	}
	fmt.Println("task", r)
	return r
}
