package main

import (
	"fmt"
	"sync"
)

type task func()
type deque struct {
	dqmutex sync.Mutex
	deque   []task
}

func (d *deque) pushFront(t task) {
	d.dqmutex.Lock()
	defer d.dqmutex.Unlock()
	d.deque = append(d.deque, t)
}

func (d *deque) pushBack(t task) {
	d.dqmutex.Lock()
	defer d.dqmutex.Unlock()
	d.deque = append([]task{t}, d.deque...)
}

func (d *deque) popFront() (task, error) {
	d.dqmutex.Lock()
	defer d.dqmutex.Unlock()
	if len(d.deque) == 0 {
		return nil, fmt.Errorf("deque is empty")
	}
	t := d.deque[0]
	d.deque = d.deque[1:]
	return t, nil
}
func (d *deque) popBack() (task, error) {
	d.dqmutex.Lock()
	defer d.dqmutex.Unlock()
	if len(d.deque) == 0 {
		return nil, fmt.Errorf("deque is empty")
	}
	t := d.deque[len(d.deque)-1]
	d.deque = d.deque[:len(d.deque)-1]
	return t, nil
}
