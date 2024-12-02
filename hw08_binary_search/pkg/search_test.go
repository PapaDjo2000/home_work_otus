package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch_BinarySearch(t *testing.T) {
	mass := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	testCases := []struct {
		name   string
		mass   []int
		target int
		want   int
	}{
		{
			name:   "1",
			mass:   mass,
			target: 4,
			want:   3,
		},
		{
			name:   "2",
			mass:   mass,
			target: 15,
			want:   -1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			want := BinarySearch(tC.target, tC.mass)
			assert.Equal(t, tC.want, want)
		})
	}
}
