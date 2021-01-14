package domap

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func Worker(m *Master, wg *sync.WaitGroup) *Result {
	defer wg.Done()
	timeout := time.After(time.Second * time.Duration(m.timeout))
	done := make(chan bool, 1)
	for {
		select {
		case <-done:
			continue
		case <-timeout:
			return nil
		case <-m.stop:
			return nil
		default:
			go func(done chan bool) {
				defer func() {
					if p := recover(); p != nil {
						err := errors.New(fmt.Sprintf("panic: %s\n", p))
						result := &Result{err: err}
						m.SetRes(result)
					}
					done <- true
				}()
				task := <-m.queue
				fRes := m.f(task.key)
				result := &Result{Res: fRes}
				m.SetRes(result)
			}(done)
		}
	}
}
