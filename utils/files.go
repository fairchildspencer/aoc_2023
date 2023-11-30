package utils

import (
	"fmt"
	"os"
)

func ReadInput(input string) string {
	b, err := os.ReadFile(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(b)
}
