package main

import "fmt"

func main() {

answer := fib(1)


/*var input uint=5
fmt.Println(fib(input))*/

fmt.Println(answer)
}

func fib(n uint) uint {
    if n < 2 {
        return n;
    }
    return fib(n-1) + fib(n-2);
}