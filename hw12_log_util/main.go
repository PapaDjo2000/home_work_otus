package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func analyzerLogFile(filename string, loglevel string) string {
	var count int

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("error open file; %v", err)
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if strings.Contains(line, loglevel) {
			count++
		}
	}

	return fmt.Sprintf("file statistics %s=%d", loglevel, count)
}

func main() {
	var filename, logLevel, outputfile string

	os.Setenv("LOG_ANALYZER_FILE", "./log.log")
	os.Setenv("LOG_ANALYZER_LEVEL", "INFO")
	os.Setenv("LOG_ANALYZER_OUTPUT", "./Output.txt")

	flag.StringVar(&filename, "file", "", "path to logfile")
	flag.StringVar(&logLevel, "level", "INFO", "Log level to analyze")
	flag.StringVar(&outputfile, "output", "", "path to output  analyze")

	flag.Parse()

	if filename == "" {
		filename = os.Getenv("LOG_ANALYZER_FILE")
	}
	if logLevel == "" {
		logLevel = os.Getenv("LOG_ANALYZER_LEVEL")
	}
	if outputfile == "" {
		outputfile = os.Getenv("LOG_ANALYZER_OUTPUT")
	}
	resultCount := analyzerLogFile(filename, logLevel)

	if outputfile == "" {
		fmt.Println(resultCount)
	} else {
		err := os.WriteFile(outputfile, []byte(resultCount), 0o600)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Statistics written to %s\n", outputfile)
	}
}
