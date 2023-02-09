package main

import "testing"

func Test_pipe(t *testing.T) {
	s, err := pipe("echo test | grep t")
	if err != nil {
		t.Error(err)
	}

	if s != "test" {
		t.Errorf("got %s, want %s", s, "test")
	}
}
