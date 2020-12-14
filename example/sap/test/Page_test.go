package test

import (
	"fmt"
	"sap/Logic"
	"testing"
)

func TestPage(t *testing.T) {
	type test struct {
		inputA int
	}
	tests := []test{
		{inputA: 10},
		{inputA: 1},
		{inputA: 15},
		{inputA: 37},
	}
	for _, tc := range tests {
		l,s:=Logic.Page(tc.inputA)
		fmt.Println(l,s)
	}
}
