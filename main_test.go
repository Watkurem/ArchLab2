package main

import (
	"testing"
	"bytes"
)

const (
	EMPTY_LINE = "\n"
	VOID = ""
)

func test_ok(str, msk, expect string, t *testing.T) {
	if expect != VOID && expect != EMPTY_LINE {
		expect += "\n"
	}

	out = new(bytes.Buffer)
	print_matches(str, msk)
	got := out.(*bytes.Buffer).String()
	if got != expect {
		t.Errorf("String was \"%s\", mask was \"%s\"." +
			"Expected \"%s\"." +
			"Got \"%s\"",
			str, msk, expect, got)
	}
}

func Test_empty_line(t *testing.T) {
	test_ok("", "", EMPTY_LINE, t)
}

func Test_exact_match_succes(t *testing.T) {
	test_ok("abcdefg", "cde", "cde", t)
}

func Test_exact_match_fail(t *testing.T) {
	test_ok("abcdefg", "fgab", VOID, t)
}

func Test_exact_match_asterisk(t *testing.T) {
	test_ok("abc", "*",
		`
a
ab
abc
b
bc
c`,
		t)
}
