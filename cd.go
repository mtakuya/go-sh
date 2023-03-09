package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"syscall"
)

func cd(s string) (string, error) {
	c := strings.Split(s, " ")
	if len(c) != 2 {
		return "", errors.New("cd error")
	}

	var path string
	if c[1] == "~" {
		path = os.Getenv("HOME")
	} else {
		path = c[1]
	}

	err := syscall.Chdir(path)
	if err != nil {
		return "", fmt.Errorf("%s %s", err, path)
	}

	d, err := os.Getwd()
	if err != nil {
		return "", err
	} else {
		return d, nil
	}
}
