package main

import (
	"fmt"
	"os"
	"strings"
)

func env() (string, error) {
	envs := os.Environ()
	var sb strings.Builder

	for _, v := range envs {
		_, err := fmt.Fprintln(&sb, v)
		if err != nil {
			return "", err
		}
	}

	return strings.TrimRight(sb.String(), "\n"), nil
}
