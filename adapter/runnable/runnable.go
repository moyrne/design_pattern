package runnable

import "log"

type (
	Runnable interface {
		Run()
	}
	Thread struct {
		runnable Runnable
	}
)

func NewThread(runnable Runnable) *Thread {
	return &Thread{runnable: runnable}
}

func (t *Thread) Start() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()
		t.runnable.Run()
	}()
}
