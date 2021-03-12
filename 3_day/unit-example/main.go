package main

var opMap map[string]func(int, int) int

func initOpMap() {
	opMap = make(map[string]func(int, int) int)
	opMap["+"] = add
	opMap["_"] = sub
	opMap["*"] = mul
	opMap["/"] = div
}
func Calculate(op string, a, b int) int {
	if v, ok := opMap[op]; ok {
		return v(a, b)
	}
	return 0
}

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func mul(a, b int) int {
	return a * b
}
func div(a, b int) int {
	return a / b
}

func main() {
	initOpMap()
	// result := Calculate("+", 1, 2)
	// print(result)
}
