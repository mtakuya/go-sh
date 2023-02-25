package main

import "testing"

func Test_tee(t *testing.T) {
	result, err := pipe("ls | tee ls.txt")
	defer rm("rm ls.txt")

	if err != nil {
		t.Error(err)
	}

	if result == "" {
		t.Errorf("got blank")
	}
}
