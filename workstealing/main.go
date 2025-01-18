package main

import "time"

func main() {
	workers := make([]*worker, 0)
	no_of_workers := 5
	for i := 0; i < no_of_workers; i++ {
		workers = append(workers, &worker{dq: &deque{}, active: true})
	}
	workerPool := &WorkerPool{workers: workers}
	for _, w := range workerPool.workers {
		w.workerPool = workerPool
	}
	workerPool.startWorkers()
	for i := 0; i < 100; i++ {
		workerPool.submitTask()
	}
	time.Sleep(100 * time.Second)
}
