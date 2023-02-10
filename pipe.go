package main

import (
	"strings"
)

func pipe(t string) (string, error) {
	var tss []string
	var result string
	var err error

	c := strings.Split(t, "|")

	for _, s := range c {
		t1 := strings.TrimLeft(s, " ")
		t2 := strings.TrimRight(t1, " ")
		tss = append(tss, t2)
	}

	for _, s := range tss {
		if result == "" {
			result, err = exec(s)
		} else {
			result, err = exec(s + " " + result)
		}
	}
	if err != nil {
		return "", err
	}
	return result, err
}
