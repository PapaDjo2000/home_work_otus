package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Analiz(t *testing.T) {
	filelog := `2025-01-19 19:53:01 INFO Starting application...
2025-01-19 19:53:03 INFO Application version 1.0.0
2025-01-19 19:53:05 WARN Unable to connect to database: Connection refused 
2025-01-19 19:53:06 ERROR Failed to load configuration file: config.json 
2025-01-19 19:53:07 FATAL Critical error occurred: Division by zero 
2025-01-19 19:53:08 INFO Application stopped
2025-01-19 19:53:07 FATAL Critical error occurred: Division by zero 
2025-01-19 19:53:08 INFO Application stopped
2025-01-19 19:53:07 FATAL Critical error occurred: Division by zero 
2025-01-19 19:53:08 INFO Application stopped`
	level := "INFO"

	err := os.WriteFile("./test.log", []byte(filelog), 0o644)
	if err != nil {
		fmt.Print(err)
	}
	want := "file statistics INFO=5"

	result := analyzerLogFile("./test.log", level)
	assert.Equal(t, want, result)
	err = os.Remove("./test.log")
	if err != nil {
		fmt.Print(err)
	}
}
