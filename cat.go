package main

import (
	"os"
	"strings"
)

func cat(s string) (string, error) {
	c := strings.Split(s, " ")
	b, err := os.ReadFile(c[1])
	if err != nil {
		return "", err
	}
	return string(b), nil
}
