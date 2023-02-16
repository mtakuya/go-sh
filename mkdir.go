package main

import (
	"errors"
	"os"
	"strings"
)

func mkdir(s string) (string, error) {
	c := strings.Split(s, " ")
	if len(c) < 2 {
		return "", errors.New("mkdir error")
	}

	err := os.Mkdir(c[1], 0755)
	if err != nil {
		return "", err
	}
	
	return "", nil
}
