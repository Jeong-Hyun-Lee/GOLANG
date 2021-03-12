package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	// add
	addResult := Calculate("+", 1, 2)
	assert.Equal(t, addResult, 3, "add should be equal")

	addNoResult := Calculate("+", 1, 4)
	assert.Equal(t, addNoResult, 3, "add should not be equal")

	//sub
	subResult := Calculate("-", 1, 2)
	assert.Equal(t, subResult, 3, "sub should be equal")

	subNoResult := Calculate("-", 1, 2)
	assert.Equal(t, subNoResult, -1, "sub should not be equal")

	// mul
	mulResult := Calculate("*", 1, 2)
	assert.Equal(t, mulResult, 2, "mul should be equal")

	mulNoResult := Calculate("*", 1, 2)
	assert.Equal(t, mulNoResult, 3, "mul should not be equal")

	// div
	divResult := Calculate("/", 5, 1)
	assert.Equal(t, divResult, 5, "div should be equal")

	divNoResult := Calculate("/", 1, 5)
	assert.Equal(t, divNoResult, 3, "div should not be equal")
}
