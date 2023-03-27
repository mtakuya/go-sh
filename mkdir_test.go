package main

import (
	"os"
	"testing"
)

func Test_mkdir(t *testing.T) {
	_, err := mkdir("mkdir mkdir_test_directory")
	defer func() {
		_, err := rm("rm mkdir_test_directory")
		if err != nil {
			t.Error(err)
		}
	}()

	if err != nil {
		t.Error(err)
	}

	fi, err := os.Stat("mkdir_test_directory")
	if err != nil {
		t.Error(err)
	}
	if fi.Name() != "mkdir_test_directory" {
		t.Errorf("got %s, want mkdir_test", fi.Name())
	}

	_, err = mkdir("mkdir")
	if err == nil {
		t.Error("mkdir error")
	}
}
