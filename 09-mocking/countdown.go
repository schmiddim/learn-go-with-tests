package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	countDownStart = 3
	finalWord      = "Go!"
)

type Sleeper interface {
	Sleep()
}
type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct {
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func Countdown(out io.Writer, sleeper Sleeper) {

	for i := countDownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)

	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}
