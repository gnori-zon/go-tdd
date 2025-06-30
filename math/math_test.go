package math

import (
	"fmt"
	"math"
	"testing"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		seconds int
		want    float64
	}{
		{seconds: 0, want: 0},
		{seconds: 30, want: math.Pi},
		{seconds: 15, want: math.Pi / 2},
		{seconds: 45, want: 3 * math.Pi / 2},
		{seconds: 60, want: 2 * math.Pi},
		{seconds: 7, want: math.Pi / 30 * 7},
	}
	for _, testCase := range cases {
		t.Run(fmt.Sprintf("%d seconds should converts in %f radians", testCase.seconds, testCase.want), func(t *testing.T) {
			got := secondsInRadians(testCase.seconds)
			if testCase.want != got {
				t.Errorf("Wanted %v radians, but got %v", testCase.want, got)
			}
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		seconds int
		point   Point
	}{
		{0, Point{0, 1}},
		{15, Point{1, 0}},
		{30, Point{0, -1}},
		{45, Point{-1, 0}},
		{60, Point{0, 1}},
	}

	for _, testCase := range cases {
		t.Run(fmt.Sprintf("%d seconds should converts in point: %v", testCase.seconds, testCase.point), func(t *testing.T) {
			got := secondHandPoint(testCase.seconds)
			if !roughlyEqualPoint(testCase.point, got) {
				t.Errorf("Wanted %v Point, but got %v", testCase.point, got)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		minutes int
		seconds int
		want    float64
	}{
		{minutes: 0, seconds: 0, want: 0},
		{minutes: 30, seconds: 0, want: math.Pi},
		{minutes: 15, seconds: 0, want: math.Pi / 2},
		{minutes: 45, seconds: 0, want: 3 * math.Pi / 2},
		{minutes: 60, seconds: 0, want: 2 * math.Pi},
		{minutes: 7, seconds: 0, want: math.Pi / 30 * 7},
		{minutes: 0, seconds: 7, want: (math.Pi / 30 * 7) / 60},
	}
	for _, testCase := range cases {
		t.Run(fmt.Sprintf("%d minutes and %d seconds should converts in %f radians", testCase.minutes, testCase.seconds, testCase.want), func(t *testing.T) {
			got := minutesInRadians(testCase.minutes, testCase.seconds)
			if !roughlyEqualFloat64(testCase.want, got) {
				t.Errorf("Wanted %v radians, but got %v", testCase.want, got)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		minutes int
		seconds int
		point   Point
	}{
		{0, 0, Point{0, 1}},
		{15, 0, Point{1, 0}},
		{30, 0, Point{0, -1}},
		{45, 0, Point{-1, 0}},
		{60, 0, Point{0, 1}},
		{15, 30, Point{0.9986295, -0.0523359}},
		{15, 60, Point{0.9945218, -0.1045284}},
		{16, 0, Point{0.9945218, -0.1045284}},
	}

	for _, testCase := range cases {
		t.Run(fmt.Sprintf("%d minutes and %d seconds should converts in point: %v", testCase.minutes, testCase.seconds, testCase.point), func(t *testing.T) {
			got := minuteHandPoint(testCase.minutes, testCase.seconds)
			if !roughlyEqualPoint(testCase.point, got) {
				t.Errorf("Wanted %v Point, but got %v", testCase.point, got)
			}
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		hours   int
		minutes int
		seconds int
		want    float64
	}{
		{hours: 0, minutes: 0, seconds: 0, want: 0},
		{hours: 3, minutes: 0, seconds: 0, want: math.Pi / 2},
		{hours: 15, minutes: 0, seconds: 0, want: math.Pi / 2},
		{hours: 6, minutes: 0, seconds: 0, want: math.Pi},
		{hours: 18, minutes: 0, seconds: 0, want: math.Pi},
		{hours: 12, minutes: 0, seconds: 0, want: 0},
		{hours: 24, minutes: 0, seconds: 0, want: 0},
		{hours: 7, minutes: 7, seconds: 0, want: 3.6774087},
		{hours: 3, minutes: 0, seconds: 7, want: 1.5709999},
	}
	for _, testCase := range cases {
		t.Run(fmt.Sprintf("%d hours and %d minutes and %d seconds should converts in %f radians", testCase.hours, testCase.minutes, testCase.seconds, testCase.want), func(t *testing.T) {
			got := hoursInRadians(testCase.hours, testCase.minutes, testCase.seconds)
			if !roughlyEqualFloat64(testCase.want, got) {
				t.Errorf("Wanted %v radians, but got %v", testCase.want, got)
			}
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		hours   int
		minutes int
		seconds int
		point   Point
	}{
		{0, 0, 0, Point{0, 1}},
		{3, 0, 0, Point{1, 0}},
		{15, 0, 0, Point{1, 0}},
		{6, 0, 0, Point{0, -1}},
		{18, 0, 0, Point{0, -1}},
		{9, 0, 0, Point{-1, 0}},
		{21, 0, 0, Point{-1, 0}},
		{12, 0, 0, Point{0, 1}},
		{24, 0, 0, Point{0, 1}},
		{15, 30, 30, Point{0.9985834, -0.0532074}},
		{16, 1, 0, Point{0.8651514, -0.5015107}},
	}

	for _, testCase := range cases {
		t.Run(fmt.Sprintf("hours %d and %d minutes and %d seconds should converts in point: %v", testCase.hours, testCase.minutes, testCase.seconds, testCase.point), func(t *testing.T) {
			got := hourHandPoint(testCase.hours, testCase.minutes, testCase.seconds)
			if !roughlyEqualPoint(testCase.point, got) {
				t.Errorf("Wanted %v Point, but got %v", testCase.point, got)
			}
		})
	}
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}
