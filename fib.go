package main

import "fmt"

func main() {
	var x uint
	for ; x < 11; x++ {
		fmt.Println(Fib(x))
	}
}

func Fib(x uint) uint {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return Fib(x-1) + Fib(x-2)
}
