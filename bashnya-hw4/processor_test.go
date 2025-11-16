package main

import (
	"bytes"
	"testing"
)

func runUniq(t *testing.T, in string, opts Options) (string, error) {
	t.Helper()
	bufIn := bytes.NewBufferString(in)
	bufOut := bytes.NewBuffer(nil)

	p := NewUniqProcessor(opts)
	err := p.Process(bufIn, bufOut)
	return bufOut.String(), err
}

func TestProcess_Default(t *testing.T) {
	input := "1\n2\n3\n3\n4\n5\n"
	want := "1\n2\n3\n4\n5\n"

	got, err := runUniq(t, input, Options{})
	if err != nil {
		t.Fatalf("не ожидали ошибку: %v", err)
	}
	if got != want {
		t.Errorf("default uniq: ожидали:\n%q\nполучили:\n%q", want, got)
	}
}

func TestProcess_CountMode(t *testing.T) {
	input := "a\na\na\nb\nc\nc\n"
	want := "3 a\n1 b\n2 c\n"

	got, err := runUniq(t, input, Options{Mode: modeCount})
	if err != nil {
		t.Fatalf("ошибка: %v", err)
	}
	if got != want {
		t.Errorf("count mode: ожидали %q, получили %q", want, got)
	}
}

func TestProcess_DupsMode(t *testing.T) {
	input := "a\na\nb\nc\nc\nc\nd\n"
	want := "a\nc\n"

	got, err := runUniq(t, input, Options{Mode: modeDups})
	if err != nil {
		t.Fatalf("ошибка: %v", err)
	}
	if got != want {
		t.Errorf("dups mode: ожидали %q, получили %q", want, got)
	}
}

func TestProcess_UniquesMode(t *testing.T) {
	input := "a\na\nb\nc\nc\nd\n"
	want := "b\nd\n"

	got, err := runUniq(t, input, Options{Mode: modeUniques})
	if err != nil {
		t.Fatalf("ошибка: %v", err)
	}
	if got != want {
		t.Errorf("uniques mode: ожидали %q, получили %q", want, got)
	}
}

func TestProcess_IgnoreCase(t *testing.T) {
	input := "A\na\nB\nb\nc\n"
	want := "A\nB\nc\n"

	got, err := runUniq(t, input, Options{IgnoreCase: true})
	if err != nil {
		t.Fatalf("ошибка: %v", err)
	}
	if got != want {
		t.Errorf("ignore-case: ожидали %q, получили %q", want, got)
	}
}

func TestProcess_SkipFields(t *testing.T) {
	input := "x1 a\nx2 a\nx3 b\n"
	want := "x1 a\nx3 b\n"

	got, err := runUniq(t, input, Options{SkipFields: 1})
	if err != nil {
		t.Fatalf("ошибка: %v", err)
	}
	if got != want {
		t.Errorf("skip-fields: ожидали %q, получили %q", want, got)
	}
}

func TestProcess_SkipChars(t *testing.T) {
	input := "abX\nabX\nabY\n"
	want := "abX\nabY\n"

	got, err := runUniq(t, input, Options{SkipChars: 2})
	if err != nil {
		t.Fatalf("ошибка: %v", err)
	}
	if got != want {
		t.Errorf("skip-chars: ожидали %q, получили %q", want, got)
	}
}

func TestProcess_SkipFieldsAndChars(t *testing.T) {
	input := "p1 aa\np2 aa\np3 ab\n"
	want := "p1 aa\np3 ab\n"

	got, err := runUniq(t, input, Options{SkipFields: 1, SkipChars: 1})
	if err != nil {
		t.Fatalf("ошибка: %v", err)
	}
	if got != want {
		t.Errorf("skip f+s: ожидали %q, получили %q", want, got)
	}
}
