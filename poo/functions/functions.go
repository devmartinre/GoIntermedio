package main

import (
	"fmt"
	"time"
)

func main() {
	// x := 5
	// y := func() int {
	// 	return x * 2
	// }()
	// z := 10
	// w := func() int {
	// 	return z * 2
	// }()
	// fmt.Println(y, z, w)
	c := make(chan int)
	go func() {
		fmt.Println("Starting function")
		time.Sleep(5 * time.Second)
		fmt.Println("End")
		c <- 1
	}()
	<-c
}
