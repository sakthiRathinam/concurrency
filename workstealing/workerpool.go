package main

import "fmt"

type WorkerPool struct {
	workers []*worker
}

func (wp *WorkerPool) startWorkers() {
	for _, w := range wp.workers {
		w.start()
	}
}

func (wp *WorkerPool) steal() task {
	for _, w := range wp.workers {
		if len(w.dq.deque) == 0 {
			continue
		}
		task, err := w.dq.popBack()
		if err != nil {
			fmt.Println(err)
			continue
		}
		return task
	}
	return nil
}
func (wp *WorkerPool) stopWorkers() {
	for _, w := range wp.workers {
		w.stop()
	}
}

type worker struct {
	dq         *deque
	active     bool
	workerPool *WorkerPool
}

func (w *worker) start() {
	go func() {
		for w.active == true {
			if len(w.dq.deque) == 0 {
				// steal other worker tasks
				task := w.workerPool.steal()
				if task != nil {
					task()
					continue
				}
			}
			task, err := w.dq.popBack()
			if err != nil {
				fmt.Println(err)
				continue
			}
			task()
		}
	}()
}

func (w *worker) stop() {
	w.active = false
	fmt.Println("worker stopped successfully")
}
