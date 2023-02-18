package main

import (
	"strings"
	"testing"
)

func Test_time(t *testing.T) {
	_, err := _time("time")
	if err == nil {
		t.Errorf("got true, want false")
	}

	result, err := _time("time ls | grep ms")
	if err != nil {
		t.Error("time error")
	}
	if !strings.Contains(result, "ms") {
		t.Errorf("%s not contains ms", result)
	}

	if !strings.Contains(result, "μs") {
		t.Errorf("%s not contains μs", result)
	}
}
