package main

import "fmt"
import "time"


    //channel
    var c chan int

func main(){
	//array
	TestArray();

	//slice
    TestSlice();

    //maps
    TestMaps();

    Convert();

    //variadic parameter
    variadic(1,2,3,4,5);

    //functions as values
    anony();

    //callback
    callback(1,rec);

    //pointer
    pointer();

    //new type
    newType();


    //goroutine
    c = make(chan int)

    go ready("Tea",2)
    go ready("Coffee",1)
    fmt.Println("I'm waiting")
    //time.Sleep(5 * time.Second)

    for{
        j := <- c
        fmt.Printf("from channel: %d\n",j)
    }
}

func  TestArray() {
	
	var arr [10]int;
	arr[0] = 42
	arr[1] = 13
	fmt.Printf("The first element is %d\n",arr[0])
}

func TestSlice() {
	var slice_1 = make([]int,10);
	var sl_l int;
	var sl_c int;

    sl_l = len(slice_1)
    sl_c = cap(slice_1)
    
    //slice_1[9] = 'a'    
    fmt.Printf("len: %d, cap :%d\n",sl_l,sl_c)

    slice_2 := append(slice_1,2,3,4)
    fmt.Printf("len: %d, cap :%d\n",len(slice_2),cap(slice_2))

    var a = [...]int{0,1,2,3,4,5,6,7}
    var s = make([]int,6)
    n1 := copy(s,a[0:])
    n2 := copy(s,s[2:])

    fmt.Printf("Slice : %d,%d,%v\n",n1,n2,slice_2)
	
}

func TestMaps() {
	weekday := map[string]int{
    	"Mon":1,"Tue":2,"Wen":3,
    	"Thi":4,"Fri":5,"Sat":6,
    	"Sun":7,
    }

    cnt := 0 
    for _,days := range weekday{
    	cnt += days
    }

    fmt.Printf("Numbers of days in a week: %d \n",cnt)

    weekday["undeclare"] = 8
    weekday["Mon"] = 9

    var value_b int
    var value_a int
    var present bool
    
    value_b,_ = weekday["Mon"]
    fmt.Printf("value of Monday: %d \n",value_b)

    
    delete(weekday,"Mon")
    value_a,present = weekday["Mon"]
    fmt.Printf("value of Monday: %d,%b \n",value_a,present)
}

func Convert(){
	s := "footbar"
	a := []byte(s)
    //reverse a
    for i,j :=0, len(a) -1; i< j ;i,j = i+1,j-1{
    	a[i],a[j] = a[j],a[i]
    }
    fmt.Printf("%s\n",string(a))
}

func rec(i int){
	if i == 10 {
		return
	}
	rec(i+1)
	fmt.Printf("%d \n",i)
}

func variadic(arg ...int){
     for _,n := range arg {
     	fmt.Printf("And the number is: %d\n",n)
     }
}

func anony(){
	a := func() {
		println("hello,anonymous function!")
	}
	a()
}

func callback(y int, f func(int)){
   f(y)
}

func pointer(){
    var p *int
    fmt.Printf("pointer: %v\n",p)

    var i int
    i = 111
    p = &i

    fmt.Printf("ponter: %v, values %d:\n",p,*p)
}

func newType(){
    type NameAge struct {
        name string
        age  int        
    }

    a := new(NameAge)
    a.name = "Peter";a.age = 42
    fmt.Printf("%v\n",a)
}

func ready(w string, sec int){
   time.Sleep(time.Duration(sec))
   fmt.Println(w,"is ready!")
   c <- 1
}