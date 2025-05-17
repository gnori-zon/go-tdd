package main

import (
	"fmt"
	"io"
	"iter"
	"os"
	"strconv"
	"time"
)

const finalWord = "GO!"
const countDownStart = 3

type Sleeper interface {
	Sleep()
}

func CountDown(writer io.Writer, sleeper Sleeper) {
	for number := range countDownFrom(countDownStart) {
		fmt.Fprintln(writer, strconv.Itoa(number))
		sleeper.Sleep()
	}
	fmt.Fprint(writer, finalWord)
}

func countDownFrom(start int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := start; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (s *ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	CountDown(os.Stdout, sleeper)
}
