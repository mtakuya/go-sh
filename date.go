package main

import (
	"time"
)

func date() (string, error) {
	t := time.Now()
	return t.Local().String(), nil
}
