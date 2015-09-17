package main

import "testing"
import "fmt"

func test(t *testing.T, v uint, expected uint) {
	r := fib(v)
	if r != expected {
		t.Error("Expected output should  be", expected, " but actual output is ", r)
	}
}

func TestFib(t *testing.T) {
	test(t, 1, 1)
	test(t, 3, 2)
	test(t, 5, 5)
	test(t, 7, 13)
	test(t, 10, 55)

	fmt.Println(fib(20))
}
