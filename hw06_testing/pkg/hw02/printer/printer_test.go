package printer

import (
	"fmt"
	"testing"

	"github.com/PapaDjo2000/home_work_otus/hw06_testing/pkg/hw02/types"
	"github.com/stretchr/testify/assert"
)

func TestPrinter_PrintStaff(t *testing.T) {
	var emp []types.Employee = []types.Employee{
		{UserID: 12, Age: 43, Name: "Djon", DepartmentID: 13245},
		{UserID: 2, Age: 4, Name: "Djo", DepartmentID: 1324}}

	for i := 0; i < len(emp); i++ {
		want := fmt.Sprintln("UID:", emp[i].UserID, " Age:", emp[i].Age, "Name:", emp[i].Name, "DepID:", emp[i].DepartmentID)
		assert.Equal(t, want, PrintStaff(emp))
	}

}
