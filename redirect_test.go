package main

import "testing"

func Test_redirect(t *testing.T) {
	_, err := redirect("ls > ls.txt")
	defer rm("rm ls.txt")
	if err != nil {
		t.Error(err)
	}

	_, err = redirect("ls | grep main > ls.txt")
	defer rm("rm ls.txt")
	if err != nil {
		t.Error(err)
	}
}
