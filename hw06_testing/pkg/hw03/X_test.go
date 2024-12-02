package packageX

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackageX_X(t *testing.T) {
	testCases := []struct {
		name  string
		index int
		x     int
		want  string
	}{
		{
			name:  "1",
			index: 1,
			x:     5,
			want:  " ",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			a := []byte(X(tC.index, tC.x))
			assert.Equal(t, tC.want, a)
		})
	}
}
