package main

import "testing"
import "time"

type testSleeppair struct {
	value  time.Duration
	result time.Duration
}

var testsSleep = []testSleeppair{
	{time.Second, time.Second},
	{2 * time.Second, 2 * time.Second},
}

func TestSleep(t *testing.T) {
	for _, pair := range testsSleep {
		v := SleepTime(pair.value).Seconds()
		s := StdSleepTime(pair.result).Seconds()
		if int(v) != int(s) {
			t.Error(
				"For", pair.value,
				"expected", pair.result,
				"got", int(v), "s",
			)
		}
	}
}
