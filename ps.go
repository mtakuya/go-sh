package main

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	"strings"
)

func ps() (string, error) {
	prs, err := process.Processes()
	if err != nil {
		return "", err
	}
	var b strings.Builder
	for _, p := range prs {
		pid := p.Pid
		ppid, err := p.Ppid()
		if err != nil {
			return "", err
		}
		name, err := p.Name()
		if err != nil {
			return "", err
		}
		_, err = fmt.Fprintf(&b, "%d %d %s\n", pid, ppid, name)
		if err != nil {
			return "", err
		}
	}
	return strings.TrimRight(b.String(), "\n"), nil
}
