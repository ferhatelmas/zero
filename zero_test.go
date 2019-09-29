package zero

import "testing"

func TestReplace(t *testing.T) {
	testCases := []struct {
		input       string
		replacement string
		expected    string
	}{
		{"Lorem​Ipsum​Dolor", "|", "Lorem|Ipsum|Dolor"},
		{"LoremIpsumDolor", "|", "LoremIpsumDolor"},
		{"Lorem​Ipsum​Dolor", "", "LoremIpsumDolor"},
		{"LoremIpsumDolor", "", "LoremIpsumDolor"},
	}
	for _, testCase := range testCases {
		got := Replace(testCase.input, testCase.replacement)
		if got != testCase.expected {
			t.Errorf("wrong replacement (%s):\ninput: %q\nexp:   %q\ngot:   %q",
				testCase.input, testCase.replacement, testCase.expected, got)
		}
	}
}

func TestHas(t *testing.T) {
	testCases := []struct {
		input string
		has   bool
	}{
		{"Lorem​Ipsum​Dolor", true},
		{"LoremIpsumDolor", false},
	}
	for _, testCase := range testCases {
		got := Has(testCase.input)
		if got != testCase.has {
			t.Errorf("%q is expected to be %T but got %T", testCase.input, testCase.has, got)
		}
	}
}
