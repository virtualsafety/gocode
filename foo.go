package main

// #include <string.h>
// void foo(char *s, int len) {
//     strncpy(s, "foo", len);
// }
import "C"

import "fmt"
import "unsafe"

func main() {
    buf := make([]byte, 256)
    C.foo((*C.char)(unsafe.Pointer(&buf[0])), C.int(len(buf)))
    fmt.Println(string(buf))
}

