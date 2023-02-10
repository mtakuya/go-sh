package main

import (
	"errors"
	"io"
	"os"
	"strings"
)

func cp(s string) (string, error) {
	c := strings.Split(s, " ")
	if len(c) < 3 {
		return "", errors.New("cp error")
	}

	src, err := os.Open(c[1])
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := os.Create(c[2])
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return "", err
	}
	return "", nil
}