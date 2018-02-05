package main

import (
	"fmt"
	"io"
)

func prompt(message string, dest interface{}, out io.Writer, in io.Reader) (int, error) {
	fmt.Fprint(out, message)

	return fmt.Fscan(in, dest)
}

func maxStringLength(strings []string) int {
	var max int

	for _, str := range strings {
		if max < len(str) {
			max = len(str)
		}
	}

	return max
}

func padRight(str, pad string, length int) string {
	if len(pad) == 0 || length <= 0 || len(str) >= length {
		return str
	}

	for {
		str += pad
		if len(str) > length {
			return str[0:length]
		}
	}
}
