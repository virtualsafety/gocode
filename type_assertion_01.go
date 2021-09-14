//https://yourbasic.org/golang/type-assertion-switch/
//https://golang.org/ref/spec#Type_assertions
//https://flaviocopes.com/go-empty-interface/
//https://yourbasic.org/golang/interfaces-explained/

package main
import ( 
    "fmt"   
)

func main() {
	var x interface{} = "foo"

	switch v := x.(type) {
	case nil:
		fmt.Println("x is nil")            // here v has type interface{}
	case int: 
		fmt.Println("x is", v)             // here v has type int
	case bool, string:
		fmt.Println("x is bool or string") // here v has type interface{}
	default:
		fmt.Println("type unknown")        // here v has type interface{}
	}
}
