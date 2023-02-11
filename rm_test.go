package main

import "testing"

func Test_rm(t *testing.T) {
	_, err := rm("rm")
	if err == nil {
		t.Errorf("got true, want false")
	}
}
