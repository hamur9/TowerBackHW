package main

import "testing"

func TestParseOptions_PositionalIO(t *testing.T) {
	args := []string{"-c", "-i", "-f", "2", "-s", "3", "in.txt", "out.txt"}
	opts, in, out, err := ParseOptions(args)
	if err != nil {
		t.Fatalf("не ожидали ошибку: %v", err)
	}
	if opts.Mode != modeCount || !opts.IgnoreCase || opts.SkipFields != 2 || opts.SkipChars != 3 {
		t.Errorf("параметры разобраны неверно: %+v", opts)
	}
	if in != "in.txt" || out != "out.txt" {
		t.Errorf("ошибка позиционных: in=%q out=%q", in, out)
	}
}

func TestParseOptions_ConflictingModes(t *testing.T) {
	// -c и -d одновременно недопустимы.
	_, _, _, err := ParseOptions([]string{"-c", "-d"})
	if err == nil {
		t.Fatalf("ожидали ошибку при конфликте режимов")
	}
}

func TestParseOptions_MissingNumberAfterF(t *testing.T) {
	_, _, _, err := ParseOptions([]string{"-f"})
	if err == nil {
		t.Fatalf("ожидали ошибку: отсутствует число после -f")
	}
}

func TestParseOptions_MissingNumberAfterS(t *testing.T) {
	_, _, _, err := ParseOptions([]string{"-s"})
	if err == nil {
		t.Fatalf("ожидали ошибку: отсутствует число после -s")
	}
}

func TestParseOptions_NegativeNumbers(t *testing.T) {
	_, _, _, err := ParseOptions([]string{"-f", "-1"})
	if err == nil {
		t.Fatalf("ожидали ошибку: -f требует неотрицательное целое")
	}
	_, _, _, err = ParseOptions([]string{"-s", "-2"})
	if err == nil {
		t.Fatalf("ожидали ошибку: -s требует неотрицательное целое")
	}
}

func TestParseOptions_TooManyPositionals(t *testing.T) {
	_, _, _, err := ParseOptions([]string{"in", "out", "extra"})
	if err == nil {
		t.Fatalf("ожидали ошибку: слишком много позиционных аргументов")
	}
}
