package figura

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFigura_AreaCircle(t *testing.T) {
	testCases := []struct {
		name   string
		radius int
		want   float64
	}{
		{
			name:   "1",
			radius: 12,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			circle := Circle{Radius: tC.radius}
			tC.want = 3.14 * (float64(tC.radius) * float64(tC.radius))
			testPl := circle.Area()

			assert.Equal(t, tC.want, testPl)
		})
	}
}

func TestFigura_AreaRectangle(t *testing.T) {
	testCases := []struct {
		name   string
		wigth  int
		height int
		want   float64
	}{
		{
			name:   "1",
			wigth:  12,
			height: 43,
		},
		{
			name:   "2",
			wigth:  43,
			height: 134,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			rectangle := Rectangle{
				Wigth:  tC.wigth,
				Height: tC.height,
			}
			tC.want = float64(tC.wigth * tC.height)
			testPl := rectangle.Area()

			assert.Equal(t, tC.want, testPl)
		})
	}
}

func TestFigura_AreaTriangle(t *testing.T) {
	testCases := []struct {
		name   string
		basis  int
		height int
		want   float64
	}{
		{
			name:   "1",
			basis:  12,
			height: 43,
		},
		{
			name:   "2",
			basis:  43,
			height: 134,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			triangle := Triangle{
				Basis:  tC.basis,
				Height: tC.height,
			}
			tC.want = float64(tC.basis * tC.height / 2)
			testPl := triangle.Area()

			assert.Equal(t, tC.want, testPl)
		})
	}
}
