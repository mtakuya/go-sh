package main

import "testing"

func Test_cat(t *testing.T) {
	_, err := cat("cat main.go")
	if err != nil {
		t.Error(err)
	}

	_, err = cat("")
	if err == nil {
		t.Error("got true, want false")
	}

	_, err = cat("cat")
	if err == nil {
		t.Error("got true, want false")
	}

	_, err = cat("cat a b")
	if err == nil {
		t.Error("got true, want false")
	}
}
