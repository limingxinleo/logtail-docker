package main

import (
	"github.com/limingxinleo/logtail-docker/log"
	"github.com/limingxinleo/logtail-docker/task"
)

func main() {
	log.InitLogger()
	task.Run()
}
