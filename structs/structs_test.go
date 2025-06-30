package structs

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("rect perimeter should return perimeter of rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0
		assertForShape(t, got, want, &rectangle)
	})
}

func TestShapeArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: &Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "circle", shape: &Circle{Radius: 10}, want: 314.1592653589793},
		{name: "triangle", shape: &Triangle{Base: 12.0, Height: 6.0}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			assertForShape(t, got, tt.want, tt.shape)
		})
	}

}

func assertForShape(t testing.TB, got, want float64, shape Shape) {
	t.Helper()
	if got != want {
		t.Errorf("got %g want %g, for shape: %#v", got, want, shape)
	}
}
