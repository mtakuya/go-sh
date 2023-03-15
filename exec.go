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

	switch c {
	case "cd":
		result, err = cd(t)
	case "ls":
		result, err = ls()
	case "cat":
		result, err = cat(t)
	case "pwd":
		result, err = pwd()
	case "ps":
		result, err = ps()
	case "free":
		result, err = free()
	case "echo":
		result, err = echo(t)
	case "grep":
		result, err = grep(t)
	case "cp":
		result, err = cp(t)
	case "rm":
		result, err = rm(t)
	case "md5sum":
		result, err = md5sum(t)
	case "history":
		result, err = history()
	case "mkdir":
		result, err = mkdir(t)
	case "time":
		result, err = _time(t)
	case "env":
		result, err = env()
	case "tee":
		result, err = tee(t)
	case "touch":
		result, err = touch(t)
	case "date":
		result, err = date()
	case "diff":
		result, err = diff(t)
	case "uniq":
		result, err = uniq(t)
	default:
		result, err = "", fmt.Errorf("command not found %s", c)
	}

	return result, err
}
