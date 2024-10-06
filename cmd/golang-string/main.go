package main

import "fmt"

func main() {
	str := "Hello 월드!"	// 한영이 섞인 문자열

	for i := 0; i < len(str); i++ {
		fmt.Printf(" 타입:%T 값:%d 문자값:%c\n", str[i], str[i], str[i])	
	}

	// []rune 으로 타입 변환 후 한 글자씩 순회하기
	arr := []rune(str)

	for i := 0; i < len(arr); i++ {
		fmt.Printf(" 타입:%T 값:%d 문자값:%c\n", arr[i], arr[i], arr[i])	
	}

	// range 키워드를 이용해 한 글자씩 순회하기
	for _, v := range str {
		fmt.Printf(" 타입:%T 값:%d 문자값:%c\n", v, v, v)	
	}
}