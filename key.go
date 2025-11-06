package main

import "strings"

type KeyExtractor struct {
	ignoreCase bool
	skipFields int
	skipChars  int
}

func NewKeyExtractor(opts Options) *KeyExtractor {
	return &KeyExtractor{
		ignoreCase: opts.IgnoreCase,
		skipFields: opts.SkipFields,
		skipChars:  opts.SkipChars,
	}
}

func (k *KeyExtractor) Extract(line string) string {
	start := k.startAfterFields(line, k.skipFields)
	if k.skipChars > 0 {
		if start+k.skipChars < len(line) {
			start += k.skipChars
		} else {
			return ""
		}
	}
	key := line[start:]
	if k.ignoreCase {
		key = strings.ToLower(key)
	}
	return key
}

func (k *KeyExtractor) startAfterFields(s string, f int) int {
	if f <= 0 {
		return 0
	}
	inField := false
	fieldsSeen := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != ' ' {
			if !inField {
				inField = true
			}
			continue
		}
		if inField {
			fieldsSeen++
			inField = false
			if fieldsSeen == f {
				j := i
				if j < len(s) && s[j] == ' ' {
					j++
				}
				return j
			}
		}
	}
	if inField {
		fieldsSeen++
		if fieldsSeen == f {
			return len(s)
		}
	}
	return len(s)
}
