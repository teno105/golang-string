package main

import (
	"fmt"
	"unsafe"
)

func main() {
    var str string = "Hello World"
	var slice []byte = []byte(str)

    fmt.Printf("str:=t%p\n", unsafe.StringData(str))
	fmt.Printf("slice:=t%p\n", unsafe.SliceData(slice))

    slice[2] = 'a'

    fmt.Println(str)
    fmt.Printf("%s\n", slice)
}