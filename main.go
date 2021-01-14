package main

import (
	"math/rand"
	"time"
	"./domap"
)

const Timeout = 10


// 会随机超时
func GetWeather(city string) string {
	time.Sleep(time.Duration(rand.Intn(10 * 1.5)))
	return city + ":OK"
}

func main()  {
	m := domap.Helper()
	cities := []string{"Shanghai", "Beijing", "Wuhan"}
	m.SetData(cities)
	m.SetFunc(GetWeather)
	m.SetCon(10)
	m.SetTimeout(Timeout)
	m.Run()
	//time.Sleep(time.Duration(rand.Intn(1)))
	m.Stop()
	for _, res := range m.GetResults() {
		println(res.Res)
	}
}