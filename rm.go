package main

import (
	"errors"
	"os"
	"strings"
)

func rm(s string) (string, error) {
	c := strings.Split(s, " ")
	if len(c) != 2 {
		return "", errors.New("rm error")
	}

	err := os.Remove(c[1])
	if err != nil {
		return "", err
	}
	return "", nil
}
