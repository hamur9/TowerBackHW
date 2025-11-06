package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	opts, inputPath, outputPath, err := ParseOptions(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		fmt.Fprintln(os.Stderr, Usage())
		os.Exit(2)
	}

	var in io.Reader = os.Stdin
	var out io.Writer = os.Stdout

	if inputPath != "" {
		f, e := os.Open(inputPath)
		if e != nil {
			log.Fatalf("не удалось открыть входной файл: %v", e)
		}
		defer f.Close()
		in = f
	}

	if outputPath != "" {
		if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
			log.Fatalf("не удалось создать директорию для выходного файла: %v", err)
		}
		f, e := os.Create(outputPath)
		if e != nil {
			log.Fatalf("не удалось создать выходной файл: %v", e)
		}
		defer f.Close()
		out = f
	}

	proc := NewUniqProcessor(opts)
	if err := proc.Process(in, out); err != nil {
		log.Fatalf("ошибка обработки: %v", err)
	}
}
