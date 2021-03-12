package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	/**
	testing methods
	FailNow()이 호출되면, 테스트 함수를 즉시 종료하고 다음 테스트 함수를 실행한다.
	Fatal()는 로그를 출력하는 걸 제외하고 FailNow 메서드와 같은 일을 한다.
	Fail()이 호출되면, 테스트가 실패하더라도 함수를 종료하지 않고 다음 코드를 계속 실행한다.
	Error()는 로그를 출력하는 걸 제외하고 Fail 메서드와 같은 일을 한다.
	Errorf() 형식화된 로그를 출력한다. Fila 메서드와 같은 일을 한다.
	Log() 테스트 로그를 출력한다.
	Logf() 형식화된 테스트 로그를 출력한다.
	Failed() 실패하더라도 레포트하지 않는다.
	*/
	result := Sum(1, 2)
	assert.Equal(t, result, 3, "they should be equal")
}
