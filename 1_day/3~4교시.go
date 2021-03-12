// You can edit this code!
// Click here and start typing.
package main

import "fmt"

// fmt.Println("") -> 로 사용
//  import ."fmt" // .이붙을 경우 Println("")으로 사용가능, package가 여러 개고 각 package끼리 의존적일 경우 .사용은 충돌을 발생시킬 수 있음

var t int = 1

func main() {
	/////////////////////////////
	const (
		n1 = 10 * iota
		n2 = 10 * iota
		n3 = 10 * iota
	)
	// print('\n' + n1, n2 ,n3)
	//b:= 5.1
	//c:=int(b)
	// print("\nc: ",c)
	// loopStart()
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)
	}

}

func loopStart() {

	j := 1
	for j < 3 {
		j++
	Loop:
		for i := 0; i < 10; i++ {
			if i > 5 {
				break Loop
			}
			print(i, "\n")
		}
	}

}
