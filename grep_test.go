package main

import "testing"

func Test_grep(t *testing.T) {
	s, err := grep("grep foo seafood")
	if err != nil {
		t.Error(err)
	}

	if s != "seafood" {
		t.Errorf("got %s, want %s", s, "seafood")
	}

	s, err = grep("grep main main.go")
	if err != nil {
		t.Error(err)
	}
	if s == "" {
		t.Error("got true, want false")
	}

	_, err = grep("grep")
	if err == nil {
		t.Errorf("got true, want false")
	}
}
