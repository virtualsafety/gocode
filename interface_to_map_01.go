//https://gist.github.com/nevzatalkan/f5c5ef66e88dd446976401967b6731e8
package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	b := []byte(`{"key":"value"}`)

	var f interface{}
	json.Unmarshal(b, &f)

	myMap := f.(map[string]interface{})

	fmt.Println(myMap["key"])
}
