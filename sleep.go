package main

import "time"

// func main(){
	
// }

func Sleep(x time.Duration) {
	select {
	case <-time.After(x):
	}
}

func SleepTime(d time.Duration) time.Duration {
	t := time.Now()
	Sleep(d)
	return time.Since(t)
}

func StdSleepTime(d time.Duration) time.Duration {
	t := time.Now()
	time.Sleep(d)
	return time.Since(t)
}
