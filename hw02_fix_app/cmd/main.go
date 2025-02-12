package main

import (
	"fmt"

	"github.com/PapaDjo2000/home_work_otus/hw02_fix_app/printer"
	"github.com/PapaDjo2000/home_work_otus/hw02_fix_app/reader"
	"github.com/PapaDjo2000/home_work_otus/hw02_fix_app/types"
)

func main() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path)

	fmt.Println(err)

	printer.PrintStaff(staff)
}
