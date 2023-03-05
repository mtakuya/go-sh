package main

import (
	"errors"
	"os"
	"strings"
)

func redirect(s string) (string, error) {
	var result string
	var err error
	var c []string

	if !strings.Contains(s, ">") {
		return "", errors.New("redirect error")
	}

	c = strings.Split(s, ">")
	if len(c) < 2 {
		return "", errors.New("redirect error")
	}

	if strings.Contains(s, "|") {
		result, err = pipe(c[0])
	} else {
		result, err = exec(c[0])
	}

	if err != nil {
		return "", err
	}

	f, err := os.Create(strings.TrimLeft(c[1], " "))
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = f.WriteString(result)
	if err != nil {
		return "", err
	}

	return "", nil
}
