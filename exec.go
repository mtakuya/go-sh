package main

import (
	"fmt"
	"strings"
)

func exec(t string) (string, error) {
	var result string
	var err error

	s := strings.Split(t, " ")
	if len(s) < 1 {
		return "", fmt.Errorf("command not found %s", t)
	}

	c := s[0]

	if c == "cd" {
		result, err = cd(t)
	} else if c == "ls" {
		result, err = ls()
	} else if c == "cat" {
		result, err = cat(t)
	} else if c == "pwd" {
		result, err = pwd()
	} else if c == "ps" {
		result, err = ps()
	} else if c == "free" {
		result, err = free()
	} else if c == "echo" {
		result, err = echo(t)
	} else if c == "grep" {
		result, err = grep(t)
	} else if c == "cp" {
		result, err = cp(t)
	} else if c == "rm" {
		result, err = rm(t)
	} else {
		result, err = "", fmt.Errorf("command not found %s", c)
	}
	return result, err
}
