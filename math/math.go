package math

import (
	"math"
)

type Point struct {
	X, Y float64
}

func hoursInRadians(hours int, minutes int, seconds int) float64 {
	return (minutesInRadians(minutes, seconds) / 60) + (math.Pi / (6 / float64(hours%12)))
}

func hourHandPoint(hours int, minutes int, seconds int) Point {
	return angleToPoint(hoursInRadians(hours, minutes, seconds))
}

func minutesInRadians(minutes int, seconds int) float64 {
	return (secondsInRadians(seconds) / 60) + (math.Pi / (30.0 / float64(minutes)))
}

func minuteHandPoint(minutes int, seconds int) Point {
	return angleToPoint(minutesInRadians(minutes, seconds))
}

func secondsInRadians(seconds int) float64 {
	return math.Pi / (30.0 / float64(seconds))
}

func secondHandPoint(seconds int) Point {
	return angleToPoint(secondsInRadians(seconds))
}

func angleToPoint(radians float64) Point {
	return Point{X: math.Sin(radians), Y: math.Cos(radians)}
}
