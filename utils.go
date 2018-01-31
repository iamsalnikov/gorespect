package main

import "fmt"

func prompt(message string, dest interface{}) (int, error) {
	fmt.Print(message)

	return fmt.Scanln(dest)
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
	for {
		str += pad
		if len(str) > length {
			return str[0:length]
		}
	}
}
