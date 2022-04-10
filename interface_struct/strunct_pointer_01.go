package main

import (
	"fmt"
	"time"
)

type Inner struct {
	addr string
}

type Xst struct {
	test int64
	name *Inner
}

func (X *Xst) somefunc1() {
	var I Inner
	I.addr = "4444"

	X.test = 20
	X.name = &I
}

func (X Xst) somefunc2() {
	for {
		fmt.Printf("in somefunc2: %v\n", *(X.name))
		time.Sleep(3 * time.Second)
	}
}

func (X *Xst) somefunc3() {
	for {
		fmt.Printf("in somefunc3: %v\n", *(X.name))
		time.Sleep(3 * time.Second)
	}
}

func main() {
	var I Inner
	I.addr = "5555"

	var X Xst

	X.test = 0
	X.name = &I

	go X.somefunc1()
	go X.somefunc2()
	go X.somefunc3()

	time.Sleep(30 * time.Second)
	fmt.Printf("%v\n", *(X.name))

}
