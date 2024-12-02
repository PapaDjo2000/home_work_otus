package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain_FixString(t *testing.T) {
	str := "!строка строки: ;структуры строки."
	want := []string{"строка", "строки", "структуры", "строки"}
	assert.Equal(t, want, FixString(str))
}

func TestMain_CountWord(t *testing.T) {
	str := "!строка строки: ;структуры строки."
	want := map[string]int{"строка": 1, "строки": 2, "структуры": 1}
	assert.Equal(t, want, CountWord(str))
}
