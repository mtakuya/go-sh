package main

import "testing"

func Test_exec(t *testing.T) {
	s, err := exec("echo test")
	if err != nil {
		t.Error(err)
	}

	if s != "test" {
		t.Errorf("got %s, want %s", s, "test")
	}

	s, err = exec("")
	if err == nil {
		t.Errorf("got true, want false")
	}
}
