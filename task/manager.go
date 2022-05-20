package task

import (
	"fmt"
	"time"
)

type Manager struct {
	Tasks map[string]*Task
}

func InitManager() *Manager {
	return &Manager{make(map[string]*Task)}
}

func Run() {
	m := InitManager()
	m.Tasks["foo"] = NewTask()

	for {
		for _, task := range m.Tasks {
			if task.IsRunning == false {
				go task.Run()
			}
		}
		<-time.After(time.Minute)
		fmt.Println("调度中")
	}
}
