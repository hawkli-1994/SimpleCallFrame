package domap

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

const CityNum = 100
const Con = 10
const DefaultTimeout = 10

var Count int32 = 0

func GetWeatherObs(city string) string {
	num := atomic.LoadInt32(&Count)
	atomic.AddInt32(&Count, 1)
	//fmt.Printf("num is %d", num)
	if num < 10 {
		time.Sleep(1000 * time.Second)
		return "timeout"
	} else {
		time.Sleep(time.Second)
		return "ok"
	}
}

func TestHeadObstruction(t *testing.T) {
	// 头部阻塞场景
	// 100个城市
	// 指定10个并发数量，然后前十个一定阻塞
	// 这时头部阻塞优化自动启动
	// 最后有9990 个成功
	m := Helper()
	cities := []string{"Shanghai", "Beijing", "Wuhan"}
	for i := 3; i < CityNum; i++ {
		cities = append(cities, fmt.Sprintf("City Num: %d", i))
	}
	m.SetData(cities)
	m.SetFunc(GetWeatherObs)
	m.SetCon(Con)
	m.SetTimeout(DefaultTimeout)
	m.Run()
	m.Stop()
	successCount := 0
	timeoutCount := 0
	for _, r := range m.GetResults() {
		if r.Res == "ok" {
			successCount += 1
		}
		if r.Res == "timeout" {
			timeoutCount += 1
		}
	}
	if successCount < 10 {
		t.Errorf("头部阻塞优化失败: 成功数量: %d", successCount)
	}
	if timeoutCount != 0 {
		t.Errorf("没有正确结束, 失败数量: %d", timeoutCount)
	}
}



//func GetWeatherTimeout(city string) string {
//	time.Sleep(time.Duration(rand.Intn(10 * 1.5)))
//	return city + ":OK"
//}

//func TestTimeout(t *testing.T) {
//	// 测试超时场景
//	m := Helper()
//	cities := []string{"Shanghai", "Beijing", "Wuhan"}
//	for i := 0; i < 1000; i++ {
//		cities = append(cities, fmt.Sprintf("City Num: %d", i))
//	}
//	m.SetData(cities)
//	m.SetFunc(GetWeatherTimeout)
//	m.SetCon(10)
//	m.SetTimeout(DefaultTimeout)
//	m.Run()
//	m.Stop()
//}
