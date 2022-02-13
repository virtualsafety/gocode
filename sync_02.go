//https://yourbasic.org/golang/mutex-explained/
//https://eli.thegreenplace.net/2018/beware-of-copying-mutexes-in-go/
//https://gobyexample.com/mutexes
//https://gobyexample.com/waitgroups

import (
  "fmt"
  "sync" 
  "time"
)

type Container struct {
  sync.Mutex                       // <-- Added a mutex
  counters map[string]int
}

func (c *Container) inc(name string) {
  c.Lock()                         // <-- Added locking of the mutex
  defer c.Unlock()
  c.counters[name]++
}



func main() {
  c := Container{counters: map[string]int{"a": 0, "b": 0}}

  var wg sync.WaitGroup
  doIncrement := func(name string, n int) {
	wg.Add(1)

    for i := 0; i < n; i++ {
      c.inc(name)
    }
	wg.Done()
  }

  go doIncrement("a", 100000)
  go doIncrement("a", 100000)
  go doIncrement("b", 100000)
  go doIncrement("b", 100000)

  // Wait a bit for the goroutines to finish
  wg.Wait()
  time.Sleep(300 * time.Millisecond)
  fmt.Println(c.counters)
}
