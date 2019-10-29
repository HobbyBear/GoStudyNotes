package masterworker

import (
	"fmt"
	"testing"
)

type Job struct {
	num int
}

func NewJob(num int) Job {
	return Job{num: num}
}

type Worker struct {
	id         int                 // workerID
	WorkerPool chan chan Job       // worker 池
	JobChannel chan Job            // worker 从jobChannel 中获取Job进行处理
	Result     map[interface{}]int // worker 将处理结果放入result
	quit       chan bool           // 停止worker信号
}

func NewWorker(workerPool chan chan Job, result map[interface{}]int, id int) Worker {
	return Worker{
		id:         id,
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		Result:     result,
		quit:       make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				x := job.num * job.num
				fmt.Println(w.id, ":", x)
				w.Result[x] = w.id
			case <-w.quit:
				return
			}
		}
	}()
}

func Test01(t *testing.T) {
	fmt.Println()
}
