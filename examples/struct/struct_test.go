package main

import "testing"

func TestObjectSizes(t *testing.T) {

	t.Run("rectangle perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		actual := Perimeter(rectangle)
		expected := 40.0

		if actual != expected {
			t.Errorf("got %.2f want %.2f", actual, expected)
		}
	})

	t.Run("rectangle area", func(t *testing.T) {
		rectangle := Rectangle{12.0, 10.0}
		actual := rectangle.Area()
		expected := 120.0

		if actual != expected {
			t.Errorf("got %.2f want %.2f", actual, expected)
		}
	})

	t.Run("circle area", func(t *testing.T) {
		circle := Circle{10}
		actual := circle.Area()
		expected := 314.1592653589793

		if actual != expected {
			t.Errorf("got %g want %g", actual, expected)
		}
	})
}
