package main

import "testing"

func Test_pwd(t *testing.T) {
	_, err := pwd()
	if err != nil {
		t.Error(err)
	}
}
