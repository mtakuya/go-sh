package main

import "testing"

func Test_cp(t *testing.T) {
	_, err := cp("cp")
	if err == nil {
		t.Error("got true, want false")
	}

	_, err = cp("cp a")
	if err == nil {
		t.Error("got true, want false")
	}

	_, err = cp("cp a b c")
	if err == nil {
		t.Error("got true, want false")
	}
}
