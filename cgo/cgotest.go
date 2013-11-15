
package cgotest

/*
#cgo LDFLAGS: -L. -lcgotest
#include <stdlib.h>
#include "cgotest_c.h"
*/
import "C"
import "fmt"
import "unsafe"


func Test(s string) {
	cs := C.CString(s)
	C.sofunc(cs)
	C.free(unsafe.Pointer(cs))
}

//export goCallback
func goCallback() {
	fmt.Printf("c say :fuck golang!\n")
	for pos, char := range "aΦx" {
      fmt.Printf("character '%c' starts at byte position %d\n", char, pos)
   }

   list := []string{"a", "b", "c", "d", "e", "f"}
   for pos, char := range list {
      fmt.Printf("character '%c' starts at byte position %d\n", char, pos)
   }

}
