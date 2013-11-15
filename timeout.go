package main
 
import (
    "fmt"
    "runtime"
    "time"
)
 
func job(ch int, timeout time.Duration) {
    t := time.NewTimer(timeout)
    defer t.Stop()
 
    flg := make(chan bool, 1)
    //执行真正的处理,这里用sleep模拟
    go func(tm int) {
        time.Sleep(time.Duration(tm) * time.Second)
        flg <- true
    }(ch)
 
    //监听通道，由于设有超时，不可能泄露
    select {
    case <-t.C:
        fmt.Println(ch, "time out")
    case <-flg:
        //正常退出
        fmt.Println("job", ch, "end")
        return
    }
}
 
func main() {
    runtime.GOMAXPROCS(4)
 
    timeout := 20 * time.Second
    //启动3个goroutine，若超时的话会自动退出
    go job(3, timeout)
    go job(200, timeout)
    go job(10, timeout)
 
    time.Sleep(60 * time.Second)
}