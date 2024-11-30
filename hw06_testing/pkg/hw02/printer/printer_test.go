package printer

import (
	"go/types"
	"testing"

	"github.com/PapaDjo2000/home_work_otus/hw06_testing/pkg/hw02/reader"
)

func TestPrinter_PrintStaff(t *testing.T) {
	js := "../../../json/data.json"
	st, _ := reader.ReadJSON(js)

	testCases := []struct {
		name string
		stf  []types.Employee
		result
	}{
		{
			name: "1",
			stf:  st,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {

		})
	}

}
