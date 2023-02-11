package main

import "testing"

func Test_echo(t *testing.T) {
	s, err := echo("echo test")
	if err != nil {
		t.Error(err)
	}

	if s != "test" {
		t.Errorf("got %s, want %s", s, "test")
	}

	_, err = echo("echo")
	if err == nil {
		t.Errorf("got true, want false")
	}
}
