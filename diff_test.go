package main

import (
	"os"
	"testing"
)

func Test_diff(t *testing.T) {
	f1, err := os.Create("test1.txt")
	defer func() {
		_, err := rm("rm test1.txt")
		if err != nil {
			t.Error(err)
		}
	}()

	defer func() {
		err := f1.Close()
		if err != nil {
			t.Error(err)
		}
	}()

	if err != nil {
		t.Error(err)
	}
	_, err = f1.WriteString("test1.txt")
	if err != nil {
		t.Error(err)
	}

	f2, err := os.Create("test2.txt")
	defer func() {
		_, err := rm("rm test2.txt")
		if err != nil {
			t.Error(err)
		}
	}()

	defer func() {
		err := f2.Close()
		if err != nil {
			t.Error(err)
		}
	}()

	if err != nil {
		t.Error(err)
	}
	_, err = f2.WriteString("test2.txt")
	if err != nil {
		t.Error(err)
	}

	result, err := diff("diff test1.txt test2.txt")
	if err != nil {
		t.Error(err)
	}

	if result == "" {
		t.Error("got true, want false")
	}

	_, err = diff("diff")
	if err == nil {
		t.Error("got true, want false")
	}
}
