// You can edit this code!
// Click here and start typing.
package main

import ."fmt"

type Vertex struct {
	X, Y int
}

func (Vert *Vertex) mula (scal int) int {
	Vert.X *= scal; Vert.Y *= scal
	return Vert.X * Vert.Y 
}

func (Vert Vertex) mulb (scal int) int {
	Vert.X *= scal; Vert.Y *= scal
	return Vert.X * Vert.Y 
}

func main() {
	
	vert:=Vertex{10, 20}
	Println("mula", vert.mula(10))
	Println("mulb", vert.mulb(10))
	
}

