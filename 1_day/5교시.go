// You can edit this code!
// Click here and start typing.
package main

import ."fmt"

func main() {
	var c []string
	var d []string
	c=[]string{"sun", "mom", "tue", "wen"} // 그냥 대괄호만 지정하면 slice / [...] 이나 숫자가 있다면 배열
	d=[]string{"aaaa","aaa","aa"} // 배열은 하나하나씩 값을 세팅하지만 slice는 한번에 값을 세팅할 수 있음
	Println(c, c[0], c[1], c[0][0])
	Println(c)
	Println(d)
	d[0] = "a4"
	Println(d)
	Println("d[0][0], d[0][1]", d[0][0], d[0][1])
	
}
