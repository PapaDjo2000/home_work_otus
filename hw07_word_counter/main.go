package main

import (
	"fmt"
	"strings"
)

func FixString(str string) []string {
	str = strings.ToLower(str)
	sli := []string{",", ".", "!", "?", "'", ":", ";", "...", "\"", "(", ")", "-"}

	str = strings.Join(strings.Fields(str), " ")

	for key := range sli {
		str = strings.ReplaceAll(str, sli[key], "")
	}

	s := strings.Split(str, " ")
	return s
}

func CountWord(str string) map[string]int {
	sli := FixString(str)

	mapa := make(map[string]int)

	for _, v := range sli {
		mapa[v]++
	}

	return mapa
}

func main() {
	str := CountWord("Строка,. которая повторяет: строку ?которая строка!")
	fmt.Println(str)
	fmt.Println(str["строка"])
}
