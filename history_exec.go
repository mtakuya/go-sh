package main

import (
	"errors"
	"strconv"
	"strings"
)

func historyExec(s string) (string, error) {
	if len(s) < 2 {
		return "", errors.New("historyExec error")
	}

	a := s[1:]
	i, err := strconv.Atoi(a)
	if err != nil {
		return "", err
	}

	if len(histories) == 0 || (len(histories)-1) < i || i < 0 {
		return "", errors.New("historyExec error")
	}

	var result string
	if strings.Contains(histories[i], ">") && !strings.HasPrefix(histories[i], "time") {
		result, err = redirect(histories[i])
	} else if strings.Contains(histories[i], "|") && !strings.HasPrefix(histories[i], "time") {
		result, err = pipe(histories[i])
	} else {
		result, err = exec(histories[i])
	}

	if err != nil {
		return "", err
	}
	return result, err
}
