package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../README.md")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	var data = make([]byte, fi.Size())

	n, err := file.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 읽기 완료")
	fmt.Println(string(data)) // 문자열로 변환하여 data의 내용 출력
}
