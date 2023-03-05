package main

import (
	"os"
	"testing"
)

func Test_touch(t *testing.T) {
	_, err := touch("touch touch_test.txt")
	if err != nil {
		t.Error(err)
	}
	defer rm("rm touch_test.txt")

	fi, err := os.Stat("touch_test.txt")
	if err != nil {
		t.Error(err)
	}

	if fi.Name() != "touch_test.txt" {
		t.Errorf("got %s, want %s", fi.Name(), "touch_test.txt")
	}
}
