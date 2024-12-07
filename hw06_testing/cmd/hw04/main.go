package main

import (
	"fmt"

	"github.com/PapaDjo2000/home_work_otus/hw04_struct_comparator/pkg/book"
)

func main() {
	newbook1 := book.NewBook("Властелин колец 1", "Толкин", 1954, 1125, 154.32)
	newbook2 := book.NewBook("Властелин колец 2", "Толкин", 1955, 954, 154.32)
	compare := book.NewCompare(book.Year)
	fmt.Println(compare.CompareBook(*newbook1, *newbook2))
	compare = book.NewCompare(book.Size)
	fmt.Println(compare.CompareBook(*newbook1, *newbook2))
	compare = book.NewCompare(book.Rate)
	fmt.Println(compare.CompareBook(*newbook1, *newbook2))
}
