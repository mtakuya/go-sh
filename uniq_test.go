package main

import (
	"testing"
)

func Test_uniq(t *testing.T) {
	_, err := uniq("uniq a\nb\nc\n")
	if err != nil {
		t.Error(err)
	}
}
