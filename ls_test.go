package main

import "testing"

func Test_ls(t *testing.T) {
	s, err := ls()
	if err != nil {
		t.Error(err)
	}
	if s == "" {
		t.Error("ls error")
	}
}
