package main

import "testing"
import "fmt"


func test(t *testing.T, actualResult float64, expectedResult float64) {

	if actualResult != expectedResult {
		t.Error("Expected result should come", expectedResult, " but in actual it is ", actualResult)
	}
}

func compute(shape Shape) float64 {
	return shape.perimeter()
}

func TestPerimeter(t *testing.T) {
	r1 := Rectangle{100, 100}
	c1 := Circle{10}

	test(t, compute(r1), 400)
	test(t, compute(c1), 62.83185307179586)
	
	r2 := Rectangle{200, 200}
	test(t, compute(r2), 800)
	
	c2 := Circle{100}
	test(t, compute(c2), 628.3185307179586)
	//test(t, compute(r2), 400)
	//fmt.Println("Rectangle Perimeter =  ", c2.perimeter())
	//fmt.Println("Circle Perimeter =", compute(c2))
}
