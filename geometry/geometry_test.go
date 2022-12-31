package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{5, 4}
	got := Perimeter(rectangle)
	want := 18.00
	if got != want {
		t.Errorf("got %.2f want %.2f ", got, want)
	}

}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{2, 3}, hasArea: 6},
		{name: "Circle", shape: Circle{2}, hasArea: 12.56},
		{name: "Trianle", shape: Triangle{12, 6}, hasArea: 36},
	}

	for _, areaTest := range areaTests {
		t.Run(areaTest.name, func(t *testing.T) {
			got := areaTest.shape.Area()
			if got != areaTest.hasArea {
				t.Errorf("%#v got %f want %f ", areaTest.shape, got, areaTest.hasArea)
			}
		})

	}

}
