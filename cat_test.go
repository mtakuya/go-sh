package main

import "testing"

func Test_cat(t *testing.T) {
	_, err := cat("cat main.go")
	if err != nil {
		t.Error(err)
	}
}
