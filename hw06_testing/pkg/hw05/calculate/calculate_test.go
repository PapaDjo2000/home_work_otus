package calculete

import (
	"testing"

	"github.com/PapaDjo2000/home_work_otus/hw06_testing/pkg/hw05/figura"
	"github.com/PapaDjo2000/home_work_otus/hw06_testing/pkg/hw05/shape"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		name    string
		Shape1  shape.Shape
		want    float64
		testerr any
	}{
		{
			name:    "circle",
			Shape1:  &figura.Circle{Radius: 24},
			want:    3.14 * (float64(24) * float64(24)),
			testerr: "432",
		},
		{
			name:    "rectangle",
			Shape1:  &figura.Rectangle{Wigth: 132, Height: 324},
			want:    float64(132 * 324),
			testerr: 12,
		},
		{
			name:    "triangle",
			Shape1:  &figura.Triangle{Basis: 12, Height: 57},
			want:    float64(12 * 57 / 2),
			testerr: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			var a float64
			test, err := CalculateArea(tC.Shape1)
			assert.NoError(t, err)
			assert.Equal(t, test, tC.want)

			test, err = CalculateArea(tC.testerr)
			assert.Error(t, err)
			assert.Equal(t, test, a)
		})
	}
}
