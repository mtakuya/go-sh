package main

import (
	"errors"
	"strings"
)

func echo(s string) (string, error) {
	c := strings.Split(s, " ")
	if len(c) < 2 {
		return "", errors.New("echo error")
	}
	return c[1], nil
}
