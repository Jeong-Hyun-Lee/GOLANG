// You can edit this code!
// Click here and start typing.
package main

import ."fmt"
import "os"

type Vertex struct {Lat, Long float64}
type Test struct {str, str2 string}
func main() {
	imap := map[int]string{
		0: "t",
		1: "t2", // comma가 있어야함
	}
	Println(imap[0], imap[1])
	Println(imap)
	
	var tmap map[string]Test
	tmap = make(map[string]Test) 
	tmap["t"] =Test{"0", "0"}
	tmap["t2"] =Test{"1", "1"}
	
	Println(tmap)
	
	// openFile('test.txt')
	println("DONE")
}

func openFile(fn string) {
	f, err := os.Open(fn)
	if err != nil { // nil은 null과 유사함
		panic(err)
	}
	defer f.Close()
}
