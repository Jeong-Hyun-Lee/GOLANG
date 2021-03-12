package greeting

import "fmt"

// Hello module
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome 5!", name)
	// pointer test
	var a int = 1
	a = 1
	var b *int = &a
	println(b)  // 0xc00008fec0
	println(*b) // 1
	return message
}
