아래는 실습 순서에 맞춰 다시 작성한 `README.md`입니다.

---

# golang-string

`golang-string`는 Golang으로 작성된 간단한 애플리케이션으로, 문자열의 사용을 익히기 위한 실습입니다.


## 프로젝트 구조

```plaintext
golang-string/
│
├── cmd/
│   └── golang-string/
│       └── main.go
│
├── go.mod
├── Makefile
└── README.md
```

## 실습 순서

### 1. 패키지 구조를 위한 디렉토리 생성

먼저 프로젝트 디렉터리를 설정하고 필요한 디렉터리들을 생성합니다.

```bash
mkdir golang-string
cd golang-string
go mod init golang-string

mkdir -p cmd/golang-string
```

### 2. 큰따옴표와 백쿼트를 사용하여 문자열 출력

`cmd/golang-string/` 디렉터리 아래에 `main.go` 파일을 생성하고,
문자열을 활용하는 코드를 작성합니다.

```go
// cmd/golang-string/main.go
package main

import "fmt"

func main() {
	// 큰따옴표로 묶으면 특수 문자가 동작합니다.
	str1 := "Hello\t'World'\n"

	// 백쿼트로 묶으면 특수 문자가 동작하지 않습니다.
	str2 := `Go is "awesome"!\nGo is simple and\t'powerful'`

	fmt.Printf(str1)
	fmt.Printf(str2)
}
```

이 코드를 통해 프로그램을 실행하면 문자열이 출력됩니다.

### 3. `Makefile` 작성

이제 프로젝트의 빌드 및 실행을 자동화하기 위한 `Makefile`을 프로젝트 루트에 작성합니다.

```makefile
# Go 관련 변수 설정
APP_NAME := golang-string
CMD_DIR := ./cmd/golang-string
BUILD_DIR := ./build

.PHONY: all clean build run test fmt vet install

all: build

# 빌드 명령어
build:
	@echo "==> Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_DIR)

# 실행 명령어
run: build
	@echo "==> Running $(APP_NAME)..."
	@$(BUILD_DIR)/$(APP_NAME)

# 코드 포맷팅
fmt:
	@echo "==> Formatting code..."
	go fmt ./...

# 코드 분석
vet:
	@echo "==> Running go vet..."
	go vet ./...

# 의존성 설치
install:
	@echo "==> Installing dependencies..."
	go mod tidy

# 테스트 실행
test:
	@echo "==> Running tests..."
	go test -v ./...

# 빌드 정리
clean:
	@echo "==> Cleaning build directory..."
	rm -rf $(BUILD_DIR)
```

`Makefile`을 이용하여 코드를 빌드하고 실행할 수 있습니다.

```bash
make run
```

이 명령어를 통해 `main.go`에서 작성한 문자열을 확인할 수 있습니다.

### 4. 큰따옴표와 백쿼트를 사용하여 여러 줄 문자열 출력

구조체 변수를 초기화 하는 방법에 대해 알아보겠습니다.

```go
// cmd/golang-string/main.go
package main

import "fmt"

func main() {
	// 큰따옴표에서 여러 줄을 표현하려면 \n을 사용
    	poet1 := "죽는 날까지 하늘을 우러러\n한 점 부끄럼이 없기를,\n잎새에 이는 바람에도\n나는 괴로워했다.\n"

	// 백쿼트에서 여러 줄 표현에 특수 문자가 필요 없음
	poet2 := `죽는 날까지 하늘을 우러러
한 점 부끄럼이 없기를,
잎새에 이는 바람에도
나는 괴로워했다.`

	fmt.Printf(poet1)
	fmt.Printf(poet2)
}
```

이제 `make run` 명령을 사용하면 문자열이 여러줄 출력됩니다.

```bash
make run
```

### 5. UTF-8 문자코드
Go는 UTF-8 문자코드를 표준 문자코드로 사용합니다.

### 6. rune 타입으로 한 문자담기
문자 하나를 표현하는 데 rune 타입을 사용합니다. UTF-8은 한 글자가 1~3바이트 크기이기 때문에 UTF-8 문자값을 가지려면 3바이트가 필요합니다. 하지만 Go 언어 기본 타입에서 3바이트 정수 타입은 제공되지 않기 때문에 rune 타입은 4바이트 정수 타입인 int32 타입의 별칭입니다.

```go
// cmd/golang-string/main.go
package main

import "fmt"

func main() {
	var char rune = '한'

	fmt.Printf("%T\n", char)	// char 타입 출력
	fmt.Println(char)		// char 값 출력
	fmt.Printf("%c\n", char)	// 문자 출력
}
```

이제 `make run` 명령을 사용하면 char의 정보가 출력됩니다.

```bash
make run
```

### 7. 문자열 순회

문자열 순회하는 방법은 3가지가 있습니다.
1. 인덱스를 사용한 바이트 단위 순회
2. []rune으로 타입 변환 후 한 글자씩 순회
3. range 키워드를 이용한 한 글자씩 순회

```go
// cmd/golang-string/main.go
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
```

이제 `make run` 명령을 사용하면 문자열 순회한 글자 정보가 출력됩니다.

```bash
make run
```

### 8. 문자열 합치기

문자열은 +과 += 연산을 사용해서 문자열을 이을 수 있습니다.

```go
// cmd/golang-string/main.go
package main

import "fmt"

func main() {
	str1 := "Hello"
	str2 := "World"

	str3 := str1 + " " + str2
	fmt.Println(str3)

	str1 += " " + str2
	fmt.Println(str1)
}
```

이제 `make run` 명령을 사용하면 문자열이 출력됩니다.

```bash
make run
```


### 9. 문자열 비교하기, 대소 비교하기

비교하기 : 연산자 ==, !=를 사용해서 문자열이 같은지 같지 않은지 비교합니다.
대소 비교하기 : >,<,>=,<= 연산자를 이용해서 비교합니다. 문자열 길이와 상관없이 앞글자부터 (같은 위치에 있는 글자끼리) 비교합니다.

```go
// cmd/golang-string/main.go
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
```

이제 `make run` 명령을 사용하면 문자열 비교 결과가 출력됩니다.

```bash
make run
```

### 10. 문자열 구조

string 타입은 Go 언어에서 제공하는 내장 타입으로 그 내부 구현은 감쳐줘 있습니다.
하지만 StringHeader 구조체로 강제 타입 변환을 하면 내부 구현을 엿볼수 있습니다.
```go
type StringHeader struct {
	Data	uintptr
	Len	int
}
```

string은 필드가 2개인 구조체입니다.
첫 번째 필드 Data는 uintptr 타입으로 문자열의 데이터가 있는 메모리 주소를 나타내는 일종의 포인터입니다.
두 번째 필드 Len은 int타입으로 문자열 길이를 나타냅니다.

```go
// cmd/golang-string/main.go
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
```

이제 `make run` 명령을 사용하면 문자열 정보가 출력됩니다.

```bash
make run
```

### 11. 문자열 불변

문자열은 불변(immutable)입니다. string 타입이 가리키는 문자열의 일부만 변경 할 수 없다는 말입니다.
아래의 예제를 통해 str과 slice가 다르 ㄴ메모리 주소를 가리키고 있는지 확인할 수 있습니다.

```go
// cmd/golang-string/main.go
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
```

이제 `make run` 명령을 사용하면 문자열이 서로 다른 주소값을 가리키는 것이 출력됩니다.

```bash
make run
```

### 12. 문자열 합산

Go 언어에서 string 타입 간 합 연산을 지원합니다. 한 연산을 하면 두 문자열이 하나로 합쳐지게 됩니다.

```go
// cmd/golang-string/main.go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var str string = "Hello"
	addr1 := unsafe.StringData(str)
	str += " World"
	addr2 := unsafe.StringData(str)
	str += " Welcom!"
	addr3 := unsafe.StringData(str)

	fmt.Println(str)
	fmt.Printf("addr1:\t%p\n", addr1)
	fmt.Printf("addr2:\t%p\n", addr2)
	fmt.Printf("addr3:\t%p\n", addr3)
}
```

이제 `make run` 명령을 사용하면 문자열의 메모리 주솟값이 다르게 출력됩니다.

```bash
make run
```

Go 언어는 기존 문자열 메모리 공간을 건드리지 않고, 새로운 메모리 공간을 만들어서 두 문자열을 합치기 때문에 string 합 연산 이후 주솟값이 변경됩니다. 따라서 문자열 불변 원칙이 준수됩니다.

### 12. strings.Builder 사용

string 합 연산을 빈번하게 하면 메모리가 낭비됩니다. 그래서 string 합 연산을 빈번하게 사용하는 경우에는 strings 패키지의 Builder를 이용해서 메모리 낭비를 줄일 수 있습니다.

```go
// cmd/golang-string/main.go
package main

import (
	"fmt"
	"strings"
)

func ToUpper1(str string) string {
	var rst string
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			rst += string('A' + (c - 'a'))	// 합 연산 사용
		} else {
			rst += string(c)
		}
	}
	return rst
}

func ToUpper2(str string) string {
	var builder strings.Builder
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a'))	// strings.Builder 사용
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}


func main() {
	var str string = "Hello World"

	fmt.Println(ToUpper1(str))
	fmt.Println(ToUpper2(str))
}
```

이제 `make run` 명령을 사용하면 문자열이 대문자로 변경되어 출력됩니다.

```bash
make run
```

ToUpper2()는 strings.Builder 객체를 이용해서 문자를 더합니다.
strings.Builder는 내부에 슬라이스를 가지고 있기 때문에 WriteRune() 메서드를 통해 문자를 더할 때 매번 메모리를 새로 생성하지 않고 기존 메모리 공간에 빈자리가 있으면 그냥 더하게 됩니다. 그래서 메모리 공간 낭비를 없앨 수 있습니다.
