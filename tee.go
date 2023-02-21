package main

import (
	"errors"
	"os"
	"strings"
)

func tee(s string) (string, error) {
	c := strings.Split(s, " ")
	if len(c) != 3 {
		return "", errors.New("tee error")
	}

	f, err := os.Create(c[1])
	defer f.Close()
	if err != nil {
		return "", err
	}

	_, err = f.WriteString(c[2])
	if err != nil {
		return "", err
	}

	return c[2], nil
}
