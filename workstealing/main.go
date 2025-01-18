package main

func main() {
	workers := make([]*worker, 0)
	no_of_workers := 5
	for i := 0; i < no_of_workers; i++ {
		workers = append(workers, &worker{dq: &deque{}, active: true})
	}
	workerPool := &WorkerPool{workers: workers}
	workerPool.startWorkers()

}
