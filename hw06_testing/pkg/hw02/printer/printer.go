package printer

import (
	"fmt"

	"github.com/PapaDjo2000/home_work_otus/hw06_testing/pkg/hw02/types"
)

func PrintStaff(staff []types.Employee) {
	for i := 0; i < len(staff); i++ {
		fmt.Println("UID:", staff[i].UserID, " Age:", staff[i].Age, "Name:", staff[i].Name, "DepID:", staff[i].DepartmentID)
	}
}
