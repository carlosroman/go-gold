package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/carlosroman/go-gold/internal/processor"
)

func main() {
	// TODO: replace with Viper (https://github.com/spf13/viper)
	args := os.Args[1:]

	r, err := os.Open(args[0])
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	if err := processFile(processor.New(), r, args[1], time.Now()); err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}
}

func processFile(csvProcessor processor.Processor, csvFile io.Reader, outputDir string, endDate time.Time) (err error) {
	return csvProcessor.Run(processor.NewReader(csvFile, endDate), outputDir)
}
