package main

import (
	"fmt"
	"os"
	"strings"
)

func ls() (string, error) {
	d, err := os.Getwd()

	if err != nil {
		return "", err
	}

	file, err := os.ReadDir(d)
	if err != nil {
		return "", err
	}

	var b strings.Builder
	for _, f := range file {
		_, err := fmt.Fprintf(&b, "%s\n", f.Name())
		if err != nil {
			return "", err
		}
	}
	return strings.TrimRight(b.String(), "\n"), nil
}
