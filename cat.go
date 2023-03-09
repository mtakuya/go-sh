package main

import (
	"errors"
	"os"
	"strings"
)

func cat(s string) (string, error) {
	c := strings.Split(s, " ")
	if len(c) != 2 {
		return "", errors.New("cat error")
	}

	b, err := os.ReadFile(c[1])
	if err != nil {
		return "", err
	}
	return string(b), nil
}
