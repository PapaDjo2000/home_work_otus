package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypes_String(t *testing.T) {
	testCases := []struct {
		name string
		emp  Employee
		want string
	}{
		{
			name: "1",
			emp:  Employee{UserID: 12, Age: 43, Name: "Djon", DepartmentID: 13245},
			want: "User ID: 12; Age: 43; Name: Djon; Department ID: 13245; ",
		},
	}
	for _, tC := range testCases {

		t.Run(tC.name, func(t *testing.T) {
			result := tC.emp.String()
			assert.Equal(t, tC.want, result)
		})
	}
}
