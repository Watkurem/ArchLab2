package main

import (
	"bytes"
	"testing"
)

const (
	EMPTY_LINE = "\n"
	VOID       = ""
)

func test_ok(str, msk, expect string, t *testing.T) {
	if expect != VOID && expect != EMPTY_LINE {
		expect += "\n"
	}

	out = new(bytes.Buffer)
	print_matches(str, msk)
	got := out.(*bytes.Buffer).String()
	if got != expect {
		t.Errorf("String was \"%s\", mask was \"%s\"."+
			"Expected \"%s\"."+
			"Got \"%s\"",
			str, msk, expect, got)
	}
}

func Test_empty_line(t *testing.T) {
	test_ok("", "", EMPTY_LINE, t)
}

func Test_empty_mask_fail(t *testing.T) {
	test_ok("abcdefg", "", EMPTY_LINE, t)
}

func Test_empty_string_fail(t *testing.T) {
	test_ok("", "abcdefg", VOID, t)
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

func Test_exact_match_asterisk_left_right_fail(t *testing.T) {
	test_ok("abcbd", "n*a", VOID, t)
}

func Test_exact_match_right_asterisk_success(t *testing.T) {
	test_ok("abdyurnjgv78lpo/", "jg*",
		`jg
jgv
jgv7
jgv78
jgv78l
jgv78lp
jgv78lpo
jgv78lpo/`,
		t)
}

func Test_exact_match_right_asterisk_fail(t *testing.T) {
	test_ok("abd", "j*", VOID, t)
}

func Test_exact_match_left_asterisk_success(t *testing.T) {
	test_ok("dfghrtyu", "*h",
		`dfgh
fgh
gh
h`,
		t)
}

func Test_exact_match_left_asterisk_fail(t *testing.T) {
	test_ok("dfghrtyu", "*j", VOID, t)
}

func Test_exact_match_two_asterisks_right_success(t *testing.T) {
	test_ok("aaaajvbn", "a*j*",
		`aaaaj
aaaajv
aaaajvb
aaaajvbn
aaaj
aaajv
aaajvb
aaajvbn
aaj
aajv
aajvb
aajvbn
aj
ajv
ajvb
ajvbn`,
		t)
}

func Test_exact_match_many_asterisks_success(t *testing.T) {
	test_ok("bbbbbcccccjjjjj", "b*c*j*",
		`bbbbbcccccj
bbbbbcccccjj
bbbbbcccccjjj
bbbbbcccccjjjj
bbbbbcccccjjjjj
bbbbcccccj
bbbbcccccjj
bbbbcccccjjj
bbbbcccccjjjj
bbbbcccccjjjjj
bbbcccccj
bbbcccccjj
bbbcccccjjj
bbbcccccjjjj
bbbcccccjjjjj
bbcccccj
bbcccccjj
bbcccccjjj
bbcccccjjjj
bbcccccjjjjj
bcccccj
bcccccjj
bcccccjjj
bcccccjjjj
bcccccjjjjj`,
		t)
}
