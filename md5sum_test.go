package main

import "testing"

func Test_md5sum(t *testing.T) {
	s, err := md5sum("md5 hoge")
	if err != nil {
		t.Error(err)
	}
	if s != "ea703e7aa1efda0064eaa507d9e8ab7e" {
		t.Errorf("got %s, want ea703e7aa1efda0064eaa507d9e8ab7e", s)
	}
}
