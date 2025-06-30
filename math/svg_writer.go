package math

import (
	"fmt"
	"io"
	"time"
)

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
	clockCentreX     = 150
	clockCentreY     = 150
)

func WriteToSvg(writer io.Writer, t time.Time) {
	io.WriteString(writer, svgStart)
	io.WriteString(writer, bezel)
	io.WriteString(writer, secondHandTag(t))
	io.WriteString(writer, minuteHandTag(t))
	io.WriteString(writer, hourHandTag(t))
	io.WriteString(writer, svgEnd)
}

func secondHandTag(t time.Time) string {
	point := pointForHand(secondHandPoint(t.Second()), secondHandLength)
	return fmt.Sprintf(`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, point.X, point.Y)
}

func minuteHandTag(t time.Time) string {
	point := pointForHand(minuteHandPoint(t.Minute(), t.Second()), minuteHandLength)
	return fmt.Sprintf(`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, point.X, point.Y)
}

func hourHandTag(t time.Time) string {
	point := pointForHand(hourHandPoint(t.Hour(), t.Minute(), t.Second()), hourHandLength)
	return fmt.Sprintf(`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, point.X, point.Y)
}

func pointForHand(point Point, handLength float64) Point {
	scaledPoint := Point{point.X * handLength, point.Y * handLength}
	flippedPoint := Point{scaledPoint.X, -scaledPoint.Y}
	offsetPoint := Point{flippedPoint.X + clockCentreX, flippedPoint.Y + clockCentreY}
	return offsetPoint
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
