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

func (wp *WorkerPool) steal() task {
	for _, w := range wp.workers {
		if len(w.dq.deque) == 0 {
			continue
		}
		task, err := w.dq.popFront()
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
func (wp *WorkerPool) submitTask() {

	random := rand.Intn(len(wp.workers))
	sleep_random := rand.Intn(len(wp.workers) * 2)
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
		for w.active == true {
			if len(w.dq.deque) == 0 {
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
