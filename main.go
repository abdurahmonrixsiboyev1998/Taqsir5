package main

import (
	"fmt"
)

func faktorial(s int, ch chan int) {
	kp := 1
	for i := 2; i <= s; i++ {
		kp *= i
	}
	ch <- kp
}

func main() {
	ch := make(chan int)
	var n int
	fmt.Println("Sonni kiriting: ")
	fmt.Scan(&n)
	go faktorial(n, ch)
	s := <-ch
	fmt.Println(s)
}
