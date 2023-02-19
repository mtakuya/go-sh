package main

import (
	"os"
	"testing"
)

func Test_mkdir(t *testing.T) {
	_, err := mkdir("mkdir mkdir_test")
	defer rm("mkdir_test")

	if err != nil {
		t.Error(err)
	}

	fi, err := os.Stat("mkdir_test")
	if err != nil {
		t.Error(err)
	}
	if fi.Name() != "mkdir_test" {
		t.Errorf("got %s, want mkdir_test", fi.Name())
	}
}
