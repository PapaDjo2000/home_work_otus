package main

import (
	"fmt"

	"github.com/PapaDjo2000/home_work_otus/hw02_fix_app/printer"
	"github.com/PapaDjo2000/home_work_otus/hw02_fix_app/reader"
	"github.com/PapaDjo2000/home_work_otus/hw02_fix_app/types"
)

func main() {
	path := "../../json/data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var staff []types.Employee

	if len(path) == 0 {
		path = "../../json/data.json"
	}

	staff, err := reader.ReadJSON(path)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(err)

	printer.PrintStaff(staff)
}
