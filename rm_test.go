package main

import (
	"os"
	"testing"
)

func Test_rm(t *testing.T) {
	_, err := rm("rm")
	if err == nil {
		t.Errorf("got true, want false")
	}

	_, err = touch("touch test.txt")
	if err != nil {
		t.Error(err)
	}

	_, err = rm("rm test.txt")
	if err != nil {
		t.Error(err)
	}

	_, err = os.Stat("test.txt")
	if err == nil {
		t.Error("rm error")
	}
}
