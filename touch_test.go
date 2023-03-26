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
	defer func() {
		_, err := rm("rm touch_test.txt")
		if err != nil {
			t.Error(err)
		}
	}()

	fi, err := os.Stat("touch_test.txt")
	if err != nil {
		t.Error(err)
	}

	if fi.Name() != "touch_test.txt" {
		t.Errorf("got %s, want %s", fi.Name(), "touch_test.txt")
	}

	_, err = touch("touch")
	if err == nil {
		t.Error("touch error")
	}
}
