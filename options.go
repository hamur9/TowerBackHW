package main

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	modeDefault = iota
	modeCount
	modeDups
	modeUniques
)

type Options struct {
	Mode       int
	IgnoreCase bool
	SkipFields int
	SkipChars  int
}

func Usage() string {
	return "Использование: uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]"
}

func ParseOptions(args []string) (Options, string, string, error) {
	var opts Options
	var inputPath, outputPath string
	i := 0
	setMode := func(m int) error {
		if opts.Mode != modeDefault {
			return errors.New("ошибка: параметры -c, -d, -u нельзя комбинировать")
		}
		opts.Mode = m
		return nil
	}

	for i < len(args) {
		arg := args[i]
		switch arg {
		case "-c":
			if err := setMode(modeCount); err != nil {
				return opts, "", "", err
			}
		case "-d":
			if err := setMode(modeDups); err != nil {
				return opts, "", "", err
			}
		case "-u":
			if err := setMode(modeUniques); err != nil {
				return opts, "", "", err
			}
		case "-i":
			opts.IgnoreCase = true
		case "-f":
			if i+1 >= len(args) {
				return opts, "", "", errors.New("ошибка: ожидается число после -f")
			}
			n, err := strconv.Atoi(args[i+1])
			if err != nil || n < 0 {
				return opts, "", "", errors.New("ошибка: -f требует неотрицательное целое")
			}
			opts.SkipFields = n
			i++
		case "-s":
			if i+1 >= len(args) {
				return opts, "", "", errors.New("ошибка: ожидается число после -s")
			}
			n, err := strconv.Atoi(args[i+1])
			if err != nil || n < 0 {
				return opts, "", "", errors.New("ошибка: -s требует неотрицательное целое")
			}
			opts.SkipChars = n
			i++
		default:
			if inputPath == "" {
				inputPath = arg
			} else if outputPath == "" {
				outputPath = arg
			} else {
				return opts, "", "", fmt.Errorf("слишком много позиционных аргументов: %q", arg)
			}
		}
		i++
	}

	return opts, inputPath, outputPath, nil
}
