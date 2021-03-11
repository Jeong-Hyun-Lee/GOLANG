// You can edit this code!
// Click here and start typing.
package main

import ."fmt"

func main() {
	data := []int{2,3,5,7,11,13}
	Println("data=",data)
	Printf("len(data) %d cap(data) %d\n", len(data), cap(data))
	data = data[1:4]
	Printf("data[1:4] len(data) %d cap(data) %d \n", len(data), cap(data))
	data = data[1:2]
	Printf("data[1:2] len(data) %d cap(data) %d \n", len(data), cap(data))
	data = data[:3]
	Printf("data[:3] len(data) %d cap(data) %d \n", len(data), cap(data))
	
	
}

