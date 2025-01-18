package main

import (
	"fmt"
	"math/rand"
	"time"
)

type WorkerPool struct {
	workers []*worker
}

func (wp *WorkerPool) startWorkers() {
	for _, w := range wp.workers {
		w.start()
	}
}

func (wp *WorkerPool) steal() (task, int) {
	fmt.Println("stealing got called", len(wp.workers))
	for index, w := range wp.workers {
		fmt.Println("stealing from worker ", index, w)
		if len(w.dq.deque) == 0 {
			continue
		}
		task, err := w.dq.popFront()
		if err != nil {
			fmt.Println(err)
			continue
		}
		return task, index
	}
	return nil, -1
}
func (wp *WorkerPool) stopWorkers() {
	for _, w := range wp.workers {
		w.stop()
	}
}
func (wp *WorkerPool) submitTask() {

	random := rand.Intn(len(wp.workers))
	sleep_random := rand.Intn(len(wp.workers) * 2)

	fmt.Printf("task assigned to worker %d \n", random)
	wp.workers[random].dq.pushBack(func() {
		time.Sleep(time.Duration(sleep_random) * time.Second)
		fmt.Println("task executed by worker ", random)
	})
}

type worker struct {
	dq         *deque
	active     bool
	workerPool *WorkerPool
}

func (w *worker) start() {
	go func() {

		time.Sleep(1 * time.Second)
		fmt.Println("worker started")
		for w.active {
			if len(w.dq.deque) == 0 {
				task, workerIndex := w.workerPool.steal()
				if task != nil {
					fmt.Printf("task stolen by worker %d \n", workerIndex)
					task()
					continue
				}
			}
			task, err := w.dq.popBack()
			if err != nil {
				fmt.Println(err)
				continue
			}
			if task == nil {
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
