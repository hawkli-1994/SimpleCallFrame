package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type KeyValue struct {
	Key   string
	Value string
}

type WorkError struct {
	Err string
}


func Worker(m *Master, wg *sync.WaitGroup) *Result {
	defer wg.Done()
	for {
		timeout := time.After(time.Second * time.Duration(m.timeout))
		done := make(chan bool, 1)
		go func(done chan bool) {
			defer func() {
				if p := recover(); p != nil {
					err := errors.New(fmt.Sprintf("panic: %s\n", p))
					result := &Result{err: err,}
					m.setRes(result)
				}
			}()
			task := <- m.queue
			fRes := m.f(task.key)
			result := &Result{res:fRes,}
			m.setRes(result)
		}(done)
		select {
		case <- done:
			continue
		case <- timeout:
			return nil
		case <- m.stop:
			return nil
		}
	}
}