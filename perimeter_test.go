package main

import "testing"
import "math"

type testPerimeterpair struct {
	shapes Shape
	result float64
}

var testsPe = []testPerimeterpair{
	{&Rectangle{1, 2, 3, 4}, 8},
	{&Rectangle{1, 3, 2, 4}, 2},
	{&Circle{0, 0, 3}, 2 * 3 * math.Pi},
	{&Circle{0, 0, 2}, 2 * 2 * math.Pi},
}

func TestPerimeter(t *testing.T) {
	for _, pair := range testsPe {
		v := totalPerimeter(pair.shapes)
		if v != pair.result {
			t.Error(
				"For", pair.shapes,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}
