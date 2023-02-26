package main

import "testing"

func Test_date(t *testing.T) {
	result, _ := date()
	if result == "" {
		t.Error("date error")
	}
}
