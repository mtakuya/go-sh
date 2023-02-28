package main

import "testing"

func Test_touch(t *testing.T) {
	_, err := touch("touch touch_test.txt")
	defer rm("rm touch_test.txt")

	if err != nil {
		t.Error(err)
	}
}
