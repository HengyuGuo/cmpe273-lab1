package main

import "fmt"

/*Defer will schedules a function call to be run after
  the function completes. Defer often used when resources
  need to be freed in some way. It will run even if a run-time
  panic occurs*/

/*Panic function will cause a run time error*/

/*Recover function can handle a run-time panic*/
func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC!")
}
