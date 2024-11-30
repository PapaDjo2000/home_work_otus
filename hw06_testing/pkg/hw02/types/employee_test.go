package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypes_String(t *testing.T) {
	var emp Employee
	testCases := []struct {
		name string
		str  string
	}{
		{
			name: "1",
			str:  emp.String(),
		},
	}
	for _, tC := range testCases {

		t.Run(tC.name, func(t *testing.T) {
			str := emp.String()
			assert.Equal(t, tC.str, str)
		})
	}
}
