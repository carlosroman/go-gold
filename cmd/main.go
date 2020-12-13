package main

import (
	"fmt"
	"os"

	"github.com/carlosroman/go-gold/internal/processor"
)

func main() {
	// TODO: replace with Viper (https://github.com/spf13/viper)
	args := os.Args[1:]
	if err := processFile(processor.New(), args[0], args[1]); err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

func processFile(csvProcessor processor.Processor, csvFile, outputDir string) (err error) {
	return csvProcessor.Run(csvFile, outputDir)
}
