package main 

import(
	"fmt"
	"reflect"
)

type  S struct {i int}
func (p *S) Get() int {return p.i}
func (p *S) Put(v int) {p.i = v}

type R struct {i int}
func (p *R) Get() int {return p.i}
func (p *R) Put(v int) {p.i = v}

type I interface{
	Get() int
	Put(int)
}

func  f(p I) {
	if t,ok := p.(I);ok{
	fmt.Println(reflect.TypeOf(p))
	p.Put(1)
	fmt.Println(t.Get())	
  }

  switch p.(type){
  	case *S: fmt.Println("type *S")
    case *R: fmt.Println("type *R") 
    default: fmt.Println("nothing")
  }
}

func g(p interface{}) int{
	return p.(I).Get()
}


func main() {
	var s S; f(&s);fmt.Println(s.Get())
	var r R; f(&r);fmt.Println(r.Get())
}

