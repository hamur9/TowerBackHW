package main

import "testing"

func TestKeyExtractor_IgnoreCase(t *testing.T) {
	opts := Options{IgnoreCase: true}
	k := NewKeyExtractor(opts)
	got := k.Extract("AbC")
	if got != "abc" {
		t.Errorf("ignoreCase: ожидали 'abc', получили %q", got)
	}
}

func TestKeyExtractor_SkipFields_Basic(t *testing.T) {
	opts := Options{SkipFields: 2}
	k := NewKeyExtractor(opts)
	got := k.Extract("f1 f2 key rest")
	if got != "key rest" {
		t.Errorf("skipFields: ожидали 'key rest', получили %q", got)
	}
}

func TestKeyExtractor_SkipChars_WithinLine(t *testing.T) {
	opts := Options{SkipChars: 3}
	k := NewKeyExtractor(opts)
	got := k.Extract("012345")
	if got != "345" {
		t.Errorf("skipChars: ожидали '345', получили %q", got)
	}
}

func TestKeyExtractor_SkipChars_BeyondLine(t *testing.T) {
	opts := Options{SkipChars: 10}
	k := NewKeyExtractor(opts)
	got := k.Extract("abc")
	if got != "" {
		t.Errorf("skipChars beyond: ожидали пустую строку, получили %q", got)
	}
}

func TestKeyExtractor_FieldsThenChars(t *testing.T) {
	opts := Options{SkipFields: 1, SkipChars: 2}
	k := NewKeyExtractor(opts)
	got := k.Extract("pfx KEY")
	if got != "Y" {
		t.Errorf("fields+chars: ожидали 'Y', получили %q", got)
	}
}
