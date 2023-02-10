package main

import (
	"os"
	"strings"
)

func rm(s string) (string, error) {
	c := strings.Split(s, " ")
	err := os.Remove(c[1])
	if err != nil {
		return "", err
	}
	return "", nil
}
