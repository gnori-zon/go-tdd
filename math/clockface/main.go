package main

import (
	"github.com/gnori-zon/go-tdd/math"
	"os"
	"time"
)

func main() {
	math.WriteToSvg(os.Stdout, time.Now())
}
