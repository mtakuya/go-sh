package main

import "testing"

func Test_pwd(t *testing.T) {
	result, err := pwd()
	if err != nil {
		t.Error(err)
	}

	if result == "" {
		t.Error("got true, want false")
	}
}
