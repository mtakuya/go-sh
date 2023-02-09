package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func grep(t string) (string, error) {
	s := strings.Split(t, " ")
	if len(s) < 3 {
		return "", errors.New("error")
	}

	if strings.Contains(t, "\n") {
		var b strings.Builder
		for _, ss := range strings.Split(t[len(s[0])+len(s[1])+2:], "\n") {
			if strings.Contains(ss, s[1]) {
				_, err := fmt.Fprintf(&b, "%s\n", ss)
				if err != nil {
					return "", err
				}
			}
		}
		return strings.TrimRight(b.String(), "\n"), nil
	} else {
		if _, err := os.Stat(s[2]); err == nil {
			b, err := os.ReadFile(s[2])
			if err != nil {
				return "", err
			}
			bs := strings.Split(string(b), "\n")
			var sb strings.Builder
			for _, bss := range bs {
				if strings.Contains(bss, s[1]) {
					_, err = fmt.Fprintf(&sb, "%s\n", bss)
					if err != nil {
						return "", err
					}
				}
			}
			return strings.TrimRight(sb.String(), "\n"), nil
		} else if errors.Is(err, os.ErrNotExist) {
			if strings.Contains(s[2], s[1]) {
				return s[2], nil
			} else {
				return "", nil
			}
		} else {
			return "", nil
		}
	}
}
