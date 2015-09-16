package main

import "testing"

type testFibpair struct {
	values uint
	result uint
}

var tests = []testFibpair{
	{0, 0},
	{1, 1},
	{2, 1},
	{5, 5},
	{10, 55},
}

func TestFib(t *testing.T) {
	for _, pair := range tests {
		v := Fib(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}
