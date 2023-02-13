package main

import (
	"fmt"
	"strings"
)

func history() (string, error) {
	var b strings.Builder
	for i, h := range histories {
		_, err := fmt.Fprintf(&b, "%d %s\n", i, h)
		if err != nil {
			return "", err
		}
	}
	return strings.TrimRight(b.String(), "\n"), nil
}
