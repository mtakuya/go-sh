package main

import (
	"errors"
	"os"
	"strings"
)

func touch(s string) (string, error) {
	c := strings.Split(s, " ")
	if len(c) != 2 {
		return "", errors.New("touch error")
	}

	f, err := os.Create(c[1])
	if err != nil {
		return "", err
	}
	defer f.Close()

	return "", nil
}
