package task

import (
	"fmt"
	"time"
)

type Task struct {
	StopChannel chan bool
	IsRunning   bool
	IsRemoved   bool
}

func NewTask() *Task {
	return &Task{make(chan bool, 1), false, false}
}

func (t *Task) Run() {
	t.IsRunning = true
	defer t.Stop()

	for {
		time.Sleep(time.Second)
		fmt.Println("Task Running")
	}
}

func (t *Task) Stop() {
	t.IsRunning = false
}
