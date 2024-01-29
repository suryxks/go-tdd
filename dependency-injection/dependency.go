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
const sleep = "sleep"
const write = "write"

type Sleeper interface {
	Sleep()
}
type SpySleeper struct {
	Calls int
}

type SpyCountdownOperations struct {
	Calls []string
}
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}
type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {

	s.Calls = append(s.Calls, write)
	return
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func CountDown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(writer, finalWord)

}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	CountDown(os.Stdout, sleeper)
}
