package main

import "fmt"

func Prompt(message string, dest interface{}) (int, error) {
	fmt.Print(message)

	return fmt.Scanln(dest)
}
