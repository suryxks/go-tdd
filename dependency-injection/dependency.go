package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func CountDown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(writer, finalWord)

}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleeper := &DefaultSleeper{}
	CountDown(os.Stdout, sleeper)
}
