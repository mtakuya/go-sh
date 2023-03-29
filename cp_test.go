package main

import "testing"

func Test_cp(t *testing.T) {
	_, err := cp("cp")
	if err == nil {
		t.Error("cp error")
	}
}
