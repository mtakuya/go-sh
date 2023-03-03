package main

import (
	"errors"
	"github.com/sergi/go-diff/diffmatchpatch"
	"os"
	"strings"
)

func diff(s string) (string, error) {
	c := strings.Split(s, " ")
	if len(c) != 3 {
		return "", errors.New("diff error")
	}

	f1 := c[1]
	f2 := c[2]

	bf1, err := os.ReadFile(f1)
	if err != nil {
		return "", err
	}

	bf2, err := os.ReadFile(f2)
	if err != nil {
		return "", err
	}

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(string(bf1), string(bf2), false)
	return dmp.DiffPrettyText(diffs), nil
}
