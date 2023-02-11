package main

import "testing"

func Test_cd(t *testing.T) {
	_, err := cd("cd ..")
	if err != nil {
		t.Error(err)
	}

	_, err = cd("cd")
	if err == nil {
		t.Errorf("got true, want false")
	}
}
