package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	testAllShapesArea := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{11, 5}, want: 55.0},
		{shape: Circle{20}, want: 1256.6370614359173},
		{shape: Triangle{12, 6}, want: 36.0},
	}

	for _, test := range testAllShapesArea {
		got := test.shape.Area()
		if got != test.want {
			t.Errorf("got %g, want %g", got, test.want)
		}
	}
}
