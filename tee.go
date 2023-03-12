package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func tee(s string) (string, error) {
	c := strings.Split(s, " ")
	if len(c) != 3 {
		return "", errors.New("tee error")
	}

	f, err := os.Create(c[1])
	if err != nil {
		return "", err
	}
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err = f.WriteString(c[2])
	if err != nil {
		return "", err
	}

	return c[2], nil
}
