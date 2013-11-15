package main

import "fmt"

type S struct{ i int }

func (p *S) Get() int  { return p.i }
func (p *S) Put(v int) { p.i = v }

type I interface {
	Get() int
	Put(int)
}

func f(p I) {
	fmt.Println(p.Get())
	//fmt.Println(p.(type))
	p.Put(1)
}

func main() {
	s := S{i: 1}
	f(&s)
}
