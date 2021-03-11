// You can edit this code!
// Click here and start typing.
package main

//import "fmt"

func add(a, b int) int {
	return a + b
}

var a, b int = 1, 2
func main() {
	// fmt.Println("Hello, 世界") 다른 패키지에서 사용가능한 함수는 대문자 사용, 소문자로 시작하면 build-in/private 함수
	// print("Hello")
	//result := add(a, b) // := 는 할당연산자로 var를 명시하지 않고 변수 선언과 초기화를 동시에 수행
	//r:= "hi" + " js"
	//d:= 1.0 - 3.0
	// fmt.Printf("result: %d , %s %f", result, r, d)
	
	var c int
	print(c)
}

// golang은 couchbase, mongodb와 잘맞음
// golang의 특징으로 method의 모음 - interface
// data를 모음 - structure
// channel은 프로세스간 통신을 위함
