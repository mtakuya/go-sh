package main

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
)

func free() (string, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Total: %v, Free:%v, UsedPercent:%f%%", v.Total, v.Free, v.UsedPercent), nil
}
