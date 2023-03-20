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

	_, err = uniq("uniq")
	if err == nil {
		t.Error("uniq error")
	}
}
