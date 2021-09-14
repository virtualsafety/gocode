//https://yourbasic.org/golang/type-assertion-switch/
//https://golang.org/ref/spec#Type_assertions
//https://flaviocopes.com/go-empty-interface/
//https://yourbasic.org/golang/interfaces-explained/

package main

import (
	"fmt"
	"strconv"
)

type MyStringer interface {
	String() string
}

type Temp int

func (t Temp) String() string {
	return strconv.Itoa(int(t)) + " °C"
}

type Point struct {
	x, y int
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func main() {
	var x interface{}

	x = 2.4
	fmt.Println(x) // 2.4

	x = Temp(24)
	fmt.Println(x) //24 °C

	x = &Point{1, 2}
	fmt.Println(x) // (1,2)
}
