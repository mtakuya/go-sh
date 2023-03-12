package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func cp(s string) (string, error) {
	c := strings.Split(s, " ")
	if len(c) != 3 {
		return "", errors.New("cp error")
	}

	src, err := os.Open(c[1])
	if err != nil {
		return "", err
	}
	defer func() {
		err := src.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	dst, err := os.Create(c[2])
	if err != nil {
		return "", err
	}
	defer func() {
		err := dst.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err = io.Copy(dst, src)
	if err != nil {
		return "", err
	}
	return "", nil
}
