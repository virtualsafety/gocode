package main

import (
   "fmt"
   "time"
   "os"
   "flag"
   "runtime/pprof"
   "log"
)


func fibonacci(n int) int{
	if(n < 2){
		return n
	}
	return fibonacci(n - 2) + fibonacci(n - 1)
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
func main() {  


	flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal(err)
        }
        pprof.StartCPUProfile(f)
        defer pprof.StopCPUProfile()
     }   

	t1 := time.Now()
	result := fibonacci(40)
	t2 := time.Now()
	fmt.Printf("%v,%v",result,t2.Sub(t1))

}
