package figura

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFigura_AreaCircle(t *testing.T) {
	testCases := []struct {
		name   string
		radius int
		result float64
	}{
		{
			name:   "1",
			radius: 12,
			result: 0.0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {

			var circle Circle = Circle{Radius: tC.radius}
			tC.result = 3.14 * (float64(tC.radius) * float64(tC.radius))
			testPl := circle.Area()

			assert.Equal(t, tC.result, testPl)
		})
	}
}

func TestFigura_AreaRectangle(t *testing.T) {
	testCases := []struct {
		name   string
		wigth  int
		height int
		result float64
	}{
		{
			name:   "1",
			wigth:  12,
			height: 43,
			result: 0.0,
		},
		{
			name:   "2",
			wigth:  43,
			height: 134,
			result: 0.0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			var rectangle Rectangle = Rectangle{Wigth: tC.wigth,
				Height: tC.height}
			tC.result = float64(tC.wigth * tC.height)
			testPl := rectangle.Area()

			assert.Equal(t, tC.result, testPl)
		})
	}
}

func TestFigura_AreaTriangle(t *testing.T) {
	testCases := []struct {
		name   string
		basis  int
		height int
		result float64
	}{
		{
			name:   "1",
			basis:  12,
			height: 43,
			result: 0.0,
		},
		{
			name:   "2",
			basis:  43,
			height: 134,
			result: 0.0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			var triangle Triangle = Triangle{Basis: tC.basis,
				Height: tC.height}
			tC.result = float64(tC.basis * tC.height / 2)
			testPl := triangle.Area()

			assert.Equal(t, tC.result, testPl)
		})
	}
}
