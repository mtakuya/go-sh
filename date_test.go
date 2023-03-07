package main

import "testing"

func Test_date(t *testing.T) {
	result, err := date()
	if err != nil {
		t.Error(err)
	}

	if result == "" {
		t.Error("got true, want false")
	}
}
