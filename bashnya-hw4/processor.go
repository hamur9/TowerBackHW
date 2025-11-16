package main

import (
	"bufio"
	"fmt"
	"io"
)

type UniqProcessor struct {
	opts Options
	key  *KeyExtractor
}

func NewUniqProcessor(opts Options) *UniqProcessor {
	return &UniqProcessor{opts: opts, key: NewKeyExtractor(opts)}
}

func (p *UniqProcessor) Process(in io.Reader, out io.Writer) error {
	sc := bufio.NewScanner(in)
	buf := make([]byte, 0, 1024*1024)
	sc.Buffer(buf, 1024*1024)

	var (
		prevKey   string
		prevLine  string
		count     int
		firstLine = true
	)

	flush := func() error {
		if firstLine {
			return nil
		}
		switch p.opts.Mode {
		case modeCount:
			_, err := fmt.Fprintf(out, "%d %s\n", count, prevLine)
			return err
		case modeDups:
			if count > 1 {
				_, err := fmt.Fprintf(out, "%s\n", prevLine)
				return err
			}
		case modeUniques:
			if count == 1 {
				_, err := fmt.Fprintf(out, "%s\n", prevLine)
				return err
			}
		default:
			_, err := fmt.Fprintf(out, "%s\n", prevLine)
			return err
		}
		return nil
	}

	for sc.Scan() {
		line := sc.Text()
		key := p.key.Extract(line)

		if firstLine {
			prevKey, prevLine, count, firstLine = key, line, 1, false
			continue
		}

		if key == prevKey {
			count++
		} else {
			if err := flush(); err != nil {
				return err
			}
			prevKey, prevLine, count = key, line, 1
		}
	}

	if err := sc.Err(); err != nil {
		return err
	}
	return flush()
}
