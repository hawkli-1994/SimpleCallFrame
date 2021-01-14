
# 需要提供一个 获取天气的 辅助方法：
* 尽可能的在参数要求的范围内提高并发， 比如con = 10, 期望线程数量尽量保持接近10
* 尽快的返回给用户，比如要求100，成功90， 失败10，一起返回即可
* 主线程退出时，要尽量优雅的通知子线程主动退出
* 还有一个整体超时，超过这个时间强行停止，有多少算多少返回
```Go
// Helpler().setData(1000城市).setFunc(获取单个城市天气).setCon(并发数).setTimeout(整体超时)
返回结果

// 如何获取天气

type Res struct {
    city string
    res string
    err error
 }

func GetWeather( city string) ( results []Res){
}

func main() {
   var timeout = 30
   var con = 10
   var city = []string { …… }
   var res []Res = Helper().set
}
```
