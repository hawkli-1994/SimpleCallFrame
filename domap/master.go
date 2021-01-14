package domap

import (
	"container/list"
	"sync"
	"time"
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
	mu          sync.Mutex // lock
	queue       chan *Task // 待执行
	stop        chan int   // 停止
	tasks       *list.List // 初始化
	results     []*Result  // 结果
	f           Handler    // handler
	con         int        // 并发数量
	timeout     int        // 全局超时时间
	tokenBucket chan int   // 令牌桶
	once        sync.Once
}

func (m *Master) initTokenBucket() {
	// 初始化令牌桶
	m.tokenBucket = make(chan int, m.con)
}

func (m *Master) addToken() {
	// 增加con个令牌
	// TODO 灵活传参
	for i := 0; i < m.con; i++ {
		m.tokenBucket <- 1
	}
}

func (m *Master) GetToken() {
	// 获取令牌
	<-m.tokenBucket
}

func (m *Master) SetData(args []string) *Master {
	m.tasks = list.New()
	for _, key := range args {
		m.tasks.PushBack(&Task{
			key: key,
		})
	}
	m.queue = make(chan *Task, m.tasks.Len())
	return m
}

func (m *Master) SetFunc(f Handler) *Master {
	m.f = f
	return m
}

func (m *Master) SetCon(con int) *Master {
	m.con = con
	return m
}

func (m *Master) SetTimeout(timeout int) *Master {
	m.timeout = timeout
	return m
}

// 检查头部阻塞
// threshold 头部阻塞判断阈值
// 头部阻塞优化，当检测到一段时间后, Result列表内没有值则往令牌桶内加入新值
// 暂不考虑令牌回收
func (m *Master) CheckHeadObstruction(starting time.Time, threshold int, start chan int) {
	for {
		select {
		case <-start:
			go func() {
				now := time.Now()
				if starting.Sub(now) >= time.Duration(threshold) {
					// TODO 头部阻塞优化
					m.addToken()
					return
				}
			}()
		case <-m.stop:
			return
		}
	}
}

func (m *Master) Run() int {
	m.stop = make(chan int)
	for e := m.tasks.Front(); e != nil; e = e.Next() {
		// 初始化任务队列
		task := e.Value.(*Task)
		task.timeout = m.timeout
		m.queue <- task
	}
	wg := sync.WaitGroup{}
	m.initTokenBucket() // 初始化 tasks 数量长度的令牌桶
	m.addToken()        // 只传递指定的con数量的令牌
	for i := 0; i < m.tasks.Len(); i++ {
		// 开启 城市任务数量个worker 但是只有令牌桶数量的worker可以执行
		wg.Add(1)
		go func() {
			defer wg.Done()
			Worker(m)
		}()
	}
	now := time.Now()
	start := make(chan int, 1)
	start <- 1
	threshold := 2 // 阈值设置
	go m.CheckHeadObstruction(now, threshold, start)
	wg.Wait()
	m.Stop()
	return 0
}

func (m *Master) SetRes(res *Result) {
	// 设置结果
	m.mu.Lock()
	defer m.mu.Unlock()
	m.results = append(m.results, res)
}

func (m *Master) closeStop() {
	m.once.Do(func() {
		close(m.stop)
	})
}

func (m *Master) Stop() {
	// 停止
	m.mu.Lock()
	defer m.mu.Unlock()
	m.closeStop()
}

func (m *Master) GetResults() []*Result {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.results
}

func Helper() *Master {
	return &Master{}
}
