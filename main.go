package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

var out io.Writer = os.Stdout

func filter_results_with_anchors(_results map[string]bool, anchors []string) []string {
	var results []string

lp:
	for res, _ := range _results {
		tmpres := res
		for _, anchor := range anchors {
			if !strings.Contains(tmpres, anchor) {
				continue lp
			} else {
				tmpres = tmpres[strings.Index(tmpres, anchor)+len(anchor):]
			}
		}
		results = append(results, res)
	}

	sort.Strings(results)
	return results
}

func trim_str_left(str *string, anchor string) bool {
	pos := strings.Index(*str, anchor)

	if len(anchor) > 0 {
		*str = (*str)[pos+len(anchor):]
	} else if len(*str) > 0 {
		*str = (*str)[1:]
	} else {
		return false
	}

	return true
}

func trim_str_right(str *string, anchor string) bool {
	pos := strings.LastIndex(*str, anchor)

	if len(anchor) > 0 {
		*str = (*str)[:pos]
	} else if len(*str) > 0 {
		*str = (*str)[:len(*str)-1]
	} else {
		return false
	}

	return true
}

func process_simple_mask(str string, msk string) bool {
	if !strings.Contains(msk, "*") {
		if strings.Contains(str, msk) {
			fmt.Fprintln(out, msk)
		}
		return false
	}
	return true
}

func prepare_string_mask_anchors(str string, msk string) (string, string, []string, string, string) {
	str = strings.TrimSpace(str)
	msk = strings.TrimSpace(msk)

	anchors := strings.Split(msk, "*")

	left_a := anchors[0]
	right_a := anchors[len(anchors)-1]

	return str, msk, anchors, left_a, right_a
}

func extract_rough_matches(str string, msk string, left_a string, right_a string) map[string]bool {
	results := make(map[string]bool)

	for strings.Contains(str, left_a) {
		left_a_pos := strings.Index(str, left_a)
		substr := str[left_a_pos+len(left_a):]

		for strings.Contains(substr, right_a) {
			right_a_pos := strings.LastIndex(substr, right_a)
			results[left_a+substr[:right_a_pos+len(right_a)]] = true

			if !trim_str_right(&substr, right_a) {
				break
			}
		}

		if !trim_str_left(&str, left_a) {
			break
		}
	}

	return results
}

func print_matches(str string, msk string) bool {
	str, msk, anchors, left_a, right_a := prepare_string_mask_anchors(str, msk)

	if !process_simple_mask(str, msk) {
		return false
	}

	results := extract_rough_matches(str, msk, left_a, right_a)

	for _, res := range filter_results_with_anchors(results, anchors) {
		fmt.Fprintln(out, res)
	}

	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var msk string
	var str string

	if len(os.Args) > 1 {
		msk = os.Args[1]
	} else {
		fmt.Fprintln(out, "Error: No mask provided")
		os.Exit(1)
	}

	for true {
		_str, err := reader.ReadString('\n')
		str += _str
		if err == io.EOF {
			break
		}
	}

	print_matches(str, msk)

	os.Exit(0)
}
