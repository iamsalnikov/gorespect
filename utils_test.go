package main

import "testing"

func TestMaxStringLength(t *testing.T) {
	testCases := []map[string]interface{} {
		{
			"length": 0,
			"strings": []string{},
		},
		{
			"length": 2,
			"strings": []string{"ab"},
		},
		{
			"length": 3,
			"strings": []string{"ab", "abc", "a"},
		},
	}

	for _, testCase := range testCases {
		data := testCase["strings"].([]string)
		expectedLength := testCase["length"].(int)

		max := maxStringLength(data)
		if max != expectedLength {
			t.Errorf("I expected to get %d but got %d", expectedLength, max)
		}
	}
}

func TestPadRight(t *testing.T) {
	testCases := []map[string]interface{} {
		{
			"src": "",
			"padSymbol": "",
			"padLength": -1,
			"out": "",
		},
		{
			"src": "a",
			"padSymbol": "",
			"padLength": 10,
			"out": "a",
		},
		{
			"src": "a",
			"padSymbol": " ",
			"padLength": 0,
			"out": "a",
		},
		{
			"src": "a",
			"padSymbol": " ",
			"padLength": -1,
			"out": "a",
		},
		{
			"src": "a",
			"padSymbol": " ",
			"padLength": 1,
			"out": "a",
		},
		{
			"src": "ab",
			"padSymbol": " ",
			"padLength": 1,
			"out": "ab",
		},
		{
			"src": "ab",
			"padSymbol": " ",
			"padLength": 4,
			"out": "ab  ",
		},
	}

	for _, testCase := range testCases {
		src := testCase["src"].(string)
		padSymbol := testCase["padSymbol"].(string)
		padLength := testCase["padLength"].(int)
		out := testCase["out"].(string)

		res := padRight(src, padSymbol, padLength)
		if res != out {
			t.Errorf("I expected to get \"%s\" but got \"%s\"", out, res)
		}
	}
}
