package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func uniq(s string) (string, error) {
	c := strings.Split(s, " ")

	if len(c) > 2 {
		return "", errors.New("uniq error")
	}

	var sb strings.Builder
	um := map[string]bool{}

	if strings.Contains(s, "\n") {
		for _, s := range strings.Split(c[1], "\n") {
			if _, ok := um[s]; !ok {
				um[s] = true
				_, err := fmt.Fprintf(&sb, "%s\n", s)
				if err != nil {
					return "", err
				}
			}
		}
	} else {
		f, err := os.Open(c[1])
		if err != nil {
			return "", err
		}

		defer func() {
			err := f.Close()
			if err != nil {
				fmt.Println(err)
			}
		}()

		s := bufio.NewScanner(f)

		for s.Scan() {
			t := s.Text()
			if _, ok := um[t]; !ok {
				if t != "" {
					um[t] = true
					_, err := fmt.Fprintf(&sb, "%s\n", t)
					if err != nil {
						return "", err
					}
				}
			}
		}
	}
	return strings.TrimRight(sb.String(), "\n"), nil
}
