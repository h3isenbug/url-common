package pool

import (
	"errors"
	"sync"
	"time"
)

type WorkerPool interface {
	AddJob(job func() error, onSuccess func()) error
	Run()
	GracefulShutdown()
}

var (
	ErrJobAddDeadlineExceeded = errors.New("job add deadline exceeded")
)

type jobItem struct {
	job       func() error
	onSuccess func()
}

type WorkerPoolV1 struct {
	distributor    chan jobItem
	workerCount    int
	jobAddDeadline time.Duration

	wg *sync.WaitGroup
}

func (pool WorkerPoolV1) GracefulShutdown() {
	close(pool.distributor)
	pool.wg.Wait()
}

func NewWorkerPoolV1(workerCount int, jobAddDeadline time.Duration) WorkerPool {
	var wg = &sync.WaitGroup{}
	wg.Add(workerCount)

	return &WorkerPoolV1{
		distributor:    make(chan jobItem, 0),
		workerCount:    workerCount,
		jobAddDeadline: jobAddDeadline,
		wg:             wg,
	}
}

func (pool WorkerPoolV1) Run() {
	for i := 0; i < pool.workerCount; i++ {
		go pool.worker()
	}

	pool.wg.Wait()
}

func (pool WorkerPoolV1) worker() {
	for jobItem := range pool.distributor {
		if err := jobItem.job(); err != nil {
			jobItem.onSuccess()
		}
	}
	pool.wg.Done()
}
func (pool WorkerPoolV1) AddJob(job func() error, onSuccess func()) error {
	select {
	case pool.distributor <- jobItem{job, onSuccess}:
		return nil
	case <-time.After(pool.jobAddDeadline):
		return ErrJobAddDeadlineExceeded
	}
}
