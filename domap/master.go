package domap

import (
	"container/list"
	"sync"
)

type Task struct {
	key     string
	timeout int
}

type Result struct {
	key string
	Res string
	err error
}

type Handler func(key string) string

type Master struct {
	mu      sync.Mutex
	queue   chan *Task // 待执行
	stop    chan int   // 停止
	tasks   *list.List // 初始化
	results []*Result  // 结果
	f       Handler
	con     int
	timeout int
}

func (m *Master) SetData(args []string) *Master {
	m.tasks = list.New()
	for _, key := range args {
		m.tasks.PushBack(&Task{
			key: key,
		})
	}
	return m
}

func (m *Master) SetFunc(f Handler) *Master {
	m.f = f
	return m
}

func (m *Master) SetCon(con int) *Master {
	m.con = con
	m.queue = make(chan *Task, con)
	m.stop = make(chan int, 0)
	return m
}

func (m *Master) SetTimeout(timeout int) *Master {
	m.timeout = timeout
	return m
}

func (m *Master) Run() int {
	for e := m.tasks.Front(); e != nil; e = e.Next() {
		task := e.Value.(*Task)
		task.timeout = m.timeout
		m.queue <- task
	}
	wg := &sync.WaitGroup{}
	for i := 0; i < m.con; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Worker(m)
		}()
	}
	wg.Wait()
	return 0
}

func (m *Master) SetRes(res *Result) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.results = append(m.results, res)
}

func (m *Master) Stop() {
	m.mu.Lock()
	defer m.mu.Unlock()
	close(m.stop)
}

func (m *Master) GetResults() []*Result {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.results
}

// Helpler().setData(1000城市).setFunc(获取单个城市天气).setCon(并发数).setTimeout(整体超时)

func Helper() *Master {
	return &Master{}
}
