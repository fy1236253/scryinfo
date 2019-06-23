package handler

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	EXECUTE_COUNT = 0
	m             *sync.Mutex
)

type Scheduler interface {
	Start()
	Stop()
	ExecuteTask()
}

type ThreadScheduler struct {
	Status  chan int
	AddTask chan int
}

func (ts *ThreadScheduler) Start() {
	runtime.GOMAXPROCS(1)
	ts.Status = make(chan int)
	ts.AddTask = make(chan int)
	go func() {
		for {
			select {
			case <-ts.AddTask:
				fmt.Printf("已经收到任务个数:%d\n", EXECUTE_COUNT)
			case <-ts.Status:
				fmt.Println("收到stop任务")
			default:
				time.Sleep(time.Second * 1)
				fmt.Println("...")

			}
		}
	}()
}

func (ts *ThreadScheduler) Stop() {
	ts.Status <- 1
}

func (ts *ThreadScheduler) ExecuteTask() {
	m = new(sync.Mutex)
	fmt.Println("新加入一条任务")
	EXECUTE_COUNT++
	ts.AddTask <- 1
}
