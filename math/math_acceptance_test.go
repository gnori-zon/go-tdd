package math

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"testing"
	"time"
)

type Svg struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func TestSecondHandWriteToSvg(t *testing.T) {
	cases := []struct {
		name       string
		time       time.Time
		secondLine Line
	}{
		{name: "midnight", time: timeOf(0, 0, 0), secondLine: Line{150, 150, 150, 60}},
		{name: "30 seconds", time: timeOf(0, 0, 30), secondLine: Line{150, 150, 150, 240}},
		{name: "45 seconds", time: timeOf(0, 0, 45), secondLine: Line{150, 150, 60, 150}},
	}
	for _, testCase := range cases {
		t.Run(fmt.Sprintf("SecondHand should return %s position", testCase.name), func(t *testing.T) {
			b := bytes.Buffer{}
			WriteToSvg(&b, testCase.time)
			svg := Svg{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(svg.Line, testCase.secondLine) {
				t.Errorf("want %+v but got %+v", testCase.secondLine, svg.Line)
			}
		})
	}
}

func TestMinuteHandWriteToSvg(t *testing.T) {
	cases := []struct {
		name       string
		time       time.Time
		minuteLine Line
	}{
		{name: "midnight", time: timeOf(0, 0, 0), minuteLine: Line{150, 150, 150, 70}},
		{name: "30 minutes", time: timeOf(0, 30, 0), minuteLine: Line{150, 150, 150, 230}},
		{name: "45 minutes", time: timeOf(0, 45, 0), minuteLine: Line{150, 150, 70, 150}},
	}
	for _, testCase := range cases {
		t.Run(fmt.Sprintf("MinuteHand should return %s position", testCase.name), func(t *testing.T) {
			b := bytes.Buffer{}
			WriteToSvg(&b, testCase.time)
			svg := Svg{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(svg.Line, testCase.minuteLine) {
				t.Errorf("want %+v but got %+v", testCase.minuteLine, svg.Line)
			}
		})
	}
}

func TestHourHandWriteToSvg(t *testing.T) {
	cases := []struct {
		name       string
		time       time.Time
		minuteLine Line
	}{
		{name: "midnight", time: timeOf(0, 0, 0), minuteLine: Line{150, 150, 150, 70}},
		{name: "6 hours", time: timeOf(6, 0, 0), minuteLine: Line{150, 150, 150, 200}},
		{name: "18 hours", time: timeOf(18, 0, 0), minuteLine: Line{150, 150, 150, 200}},
		{name: "9 hours", time: timeOf(9, 0, 0), minuteLine: Line{150, 150, 100, 150}},
		{name: "21 hours", time: timeOf(21, 0, 0), minuteLine: Line{150, 150, 100, 150}},
	}
	for _, testCase := range cases {
		t.Run(fmt.Sprintf("HourHand should return %s position", testCase.name), func(t *testing.T) {
			b := bytes.Buffer{}
			WriteToSvg(&b, testCase.time)
			svg := Svg{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(svg.Line, testCase.minuteLine) {
				t.Errorf("want %+v but got %+v", testCase.minuteLine, svg.Line)
			}
		})
	}
}

func timeOf(hours int, minutes int, seconds int) time.Time {
	return time.Date(1337, time.January, 1, hours, minutes, seconds, 0, time.UTC)
}

func containsLine(lines []Line, wantLine Line) bool {
	for _, line := range lines {
		if line == wantLine {
			return true
		}
	}
	return false
}
