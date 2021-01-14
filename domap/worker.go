package domap

import (
	"errors"
	"fmt"
	"time"
)

func Worker(m *Master) *Result {
	isTimeout := time.After(time.Second * time.Duration(m.timeout))
	done := make(chan bool, 1)
	done<-true
	for {
		select {
		case <-done:
			go func(done chan bool) {
				m.GetToken()
				result := &Result{}
				defer func() {
					if p := recover(); p != nil {
						err := errors.New(fmt.Sprintf("panic: %s\n", p))
						result.err = err
					}
					m.SetRes(result)
					done <- true
				}()
				task := <-m.queue
				result.key = task.key
				fRes := m.f(task.key)
				result.Res = fRes
			}(done)
		case <-isTimeout:
			return nil
		case <-m.stop:
			return nil
		default:
			continue
		}
	}
}
