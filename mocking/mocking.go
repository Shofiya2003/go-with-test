package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

var finalWord string = "Go!"

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleeper := ConfigurableSleeper{2 * time.Second, time.Sleep}
	Countdown(os.Stdout, &sleeper)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i >= 1; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprintf(out, finalWord)
}
