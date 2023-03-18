package main

import (
	"testing"
)

func Test_uniq(t *testing.T) {
	result, err := uniq("uniq a\nb\nc\nc")
	if err != nil {
		t.Error(err)
	}

	if result != "a\nb\nc" {
		t.Error("uniq error")
	}

	result, err = uniq("uniq main.go")
	if result != "" {
		t.Error("uniq error")
	}
}
