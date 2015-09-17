
package main

import (
	"fmt"
	"time"
	)

func main() {
go f(0)
var input string
fmt.Scanln(&input)

			}

func f(n int) {
				for i := 0; i < 10; i++ {
					fmt.Println(n, ":", i)
					Sleep(1000)
				}						
}

func Sleep(x int) {
					<-time.After(time.Millisecond * time.Duration(x))
}


