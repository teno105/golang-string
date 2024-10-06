package main

import "fmt"

func main() {
	// 큰따옴표에서 려러 줄을 표현하려면 \n을 사용
    poet1 := "죽는 날까지 하늘을 우러러\n한 점 부끄럼이 없기를,\n잎새에 이는 바람에도\n나는 괴로워했다.\n"

	// 백쿼트에서 려러 줄 표현에 특수 문자가 필요 없음
	poet2 := `죽는 날까지 하늘을 우러러
한 점 부끄럼이 없기를,
잎새에 이는 바람에도
나는 괴로워했다.`

    fmt.Printf(poet1)
	fmt.Printf(poet2)
}