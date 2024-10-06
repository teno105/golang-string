package main

import "fmt"

func main() {
    str1 := "Hello"
	str2 := "Hell"
	str3 := "Hello"

	fmt.Printf("%s == %s : %v\n", str1, str2, str1 == str2)
	fmt.Printf("%s != %s : %v\n", str1, str2, str1 != str2)
	fmt.Printf("%s == %s : %v\n", str1, str3, str1 == str2)
	fmt.Printf("%s != %s : %v\n", str1, str3, str1 != str3)

	str4 := "BBB"
	str5 := "aaaaAAA"
	str6 := "BBAD"
	str7 := "ZZZ"

	fmt.Printf("%s > %s : %v\n", str4, str5, str4 > str5)
	fmt.Printf("%s < %s : %v\n", str4, str6, str4 < str6)
	fmt.Printf("%s <= %s : %v\n", str4, str7, str4 <= str7)
}