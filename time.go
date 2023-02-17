package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func _time(s string) (string, error) {
	c := strings.Split(s, " ")
	if len(c) < 2 {
		return "", errors.New("time error")
	}

	var result string
	var err error

	now := time.Now()
	if strings.Contains(s, "|") {
		result, err = pipe(s[5:])
	} else {
		result, err = exec(s[5:])
	}
	after := time.Since(now)
	afterMillis := after.Milliseconds()
	afterMicros := after.Microseconds()
	result = fmt.Sprintf("%s\n%s %dms %dμs", result, s[5:], afterMillis, afterMicros)
	return result, err
}
