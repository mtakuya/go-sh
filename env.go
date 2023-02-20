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
		fmt.Fprintln(&sb, v)
	}

	return strings.TrimRight(sb.String(), "\n"), nil
}
