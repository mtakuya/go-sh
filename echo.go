package main

import "strings"

func echo(t string) (string, error) {
	s := strings.Split(t, " ")
	return s[1], nil
}
