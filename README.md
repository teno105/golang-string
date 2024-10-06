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

### 2. 포인터 변수 선언

`cmd/golang-string/` 디렉터리 아래에 `main.go` 파일을 생성하고,
struct 를 선언 및 활용하는 코드를 작성합니다.

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

이 코드를 통해 프로그램을 실행하면 구조체의 정보가 출력됩니다.

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

이 명령어를 통해 `main.go`에서 작성한 struct 정보를 확인할 수 있습니다.

### 4. 포인터 변숫값 비교하기

구조체 변수를 초기화 하는 방법에 대해 알아보겠습니다.

```go
// cmd/golang-string/main.go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

이제 `make run` 명령을 사용하면 각 house의 정보가 출력됩니다.

```bash
make run
```

### 5. UTF-8 문자코드


### 6. rune 타입으로 한 문자담기


### 7. 문자열 순회


### 8. 문자열 합치기


### 9. 문자열 비교하기, 대소 비교하기


### 10. 문자열 구조


### 11. 문자열 합산

