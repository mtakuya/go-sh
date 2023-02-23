package main

import (
	"errors"
	"fmt"
	"strings"
)

func exec(t string) (string, error) {
	var result string
	var err error

	s := strings.Split(t, " ")
	if len(s) == 0 {
		return "", errors.New("exec error")
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
	} else if c == "md5sum" {
		result, err = md5sum(t)
	} else if c == "history" {
		result, err = history()
	} else if c == "mkdir" {
		result, err = mkdir(t)
	} else if c == "time" {
		result, err = _time(t)
	} else if c == "env" {
		result, err = env()
	} else if c == "tee" {
		result, err = tee(t)
	} else if c == "touch" {
		result, err = touch(t)
	} else {
		result, err = "", fmt.Errorf("command not found %s", c)
	}
	return result, err
}
