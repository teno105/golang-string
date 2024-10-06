package main

import (
	"fmt"
	"unsafe"
)

type StringHeader struct {
	Data	uintptr
	Len	int
}

func main() {
    str1 := "안녕하세요. 한글 문자열입니다."
	str2 := str1

	fmt.Printf("str1")
	fmt.Printf("\n")
	fmt.Printf("str2")

	stringHeader1 := (*StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*StringHeader)(unsafe.Pointer(&str2))	

	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)
}