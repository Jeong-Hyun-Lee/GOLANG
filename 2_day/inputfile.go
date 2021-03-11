package main

import "os"

func main() {
	if len(os.Args) == 0 {
		return
	}
	// println(os.Args[0])
	// println(os.Args[1])
	// println(os.Args[2])
	// println(os.Args[3])
	i := 0
	for i < len(os.Args) {
		i++
		println(os.Args[i])
	}
}
