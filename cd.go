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
	if len(c) < 2 {
		return "", errors.New("cd error")
	}

	err := syscall.Chdir(c[1])
	if err != nil {
		return "", fmt.Errorf("%s %s", err, c[1])
	}

	d, err := os.Getwd()
	if err != nil {
		return "", err
	} else {
		return d, nil
	}
}
