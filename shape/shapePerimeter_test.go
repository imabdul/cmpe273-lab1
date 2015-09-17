package main

import "testing"

func test(t *testing.T, actualResult float64, expectedResult float64) {

	if actualResult != expectedResult {
		t.Error("Expected result should come", expectedResult, " but in actual it is ", actualResult)
	}
}

func compute(shape Shape) float64 {
	return shape.perimeter()
}

func TestPerimeter(t *testing.T) {
	r1 := Rectangle{200, 100}
	c1 := Circle{10}

	test(t, compute(r1), 600)
	test(t, compute(c1), 62.83185307179586)
	
	r2 := Rectangle{200, 300}
	c2 := Circle{100}
	
	test(t, compute(r2), 1000)
	test(t, compute(c2), 628.3185307179586)

}
