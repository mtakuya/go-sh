package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func grep(s string) (string, error) {
	c := strings.Split(s, " ")
	if len(c) < 3 {
		return "", errors.New("grep error")
	}

	if strings.Contains(s, "\n") {
		var b strings.Builder
		for _, ss := range strings.Split(s[len(c[0])+len(c[1])+2:], "\n") {
			if strings.Contains(ss, c[1]) {
				_, err := fmt.Fprintf(&b, "%s\n", ss)
				if err != nil {
					return "", err
				}
			}
		}
		return strings.TrimRight(b.String(), "\n"), nil
	} else {

		if _, err := os.Stat(c[2]); err == nil {
			b, err := os.ReadFile(c[2])
			if err != nil {
				return "", err
			}
			bs := strings.Split(string(b), "\n")
			var sb strings.Builder
			for _, bss := range bs {
				if strings.Contains(bss, c[1]) {
					_, err = fmt.Fprintf(&sb, "%s\n", bss)
					if err != nil {
						return "", err
					}
				}
			}
			return strings.TrimRight(sb.String(), "\n"), nil
		} else if errors.Is(err, os.ErrNotExist) {
			if strings.Contains(c[2], c[1]) {
				return c[2], nil
			} else {
				return "", nil
			}
		} else {
			return "", err
		}
	}
}
