package shapes

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("%#v got %.2f, want %.2f", shape, got, want)
		}
	}
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, want: 72.0},
		{name: "Circle", shape: Circle{10}, want: math.Pi * 100.0},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36.0},
	}
	for _, testCase := range areaTests {
		t.Run(testCase.name, func(t *testing.T) {
			checkArea(t, testCase.shape, testCase.want)
		})
	}
}
