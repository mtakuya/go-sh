package main

import (
	"bufio"
	"fmt"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
	"os"
	"strings"
	"syscall"
)

func main() {
	loop()
}

func loop() {
	for {
		fmt.Print("> ")
		s := bufio.NewScanner(os.Stdin)
		if s.Scan() {
			t := s.Text()
			if t == "exit" {
				break
			} else if strings.Contains(t, "|") {
				result, err := pipe(t)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				} else {
					fmt.Println(result)
				}
			} else {
				result, err := exec(t)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				} else {
					fmt.Println(result)
				}
			}
		}
	}
}

func pipe(t string) (string, error) {
	ss := strings.Split(t, "|")
	var tss []string
	var result string
	var err error

	for _, s := range ss {
		t1 := strings.TrimLeft(s, " ")
		t2 := strings.TrimRight(t1, " ")
		tss = append(tss, t2)
	}

	for _, s := range tss {
		if result == "" {
			result, err = exec(s)
		} else {
			result, err = exec(s + " " + result)
		}
	}
	if err != nil {
		return "", err
	}
	return result, err
}

func exec(t string) (string, error) {
	var result string
	var err error
	s := strings.Split(t, " ")
	c := s[0]

	if c == "cd" {
		result, err = cd(t)
	} else if c == "ls" {
		result, err = ls()
	} else if c == "sl" {
		result, err = sl()
	} else if c == "cat" {
		result, err = cat(t)
	} else if c == "pwd" {
		result, err = pwd()
	} else if c == "ps" {
		result, err = ps()
	} else if c == "free" {
		result, err = free()
	} else if c == "echo" {
		result, err = echo(t)
	} else if c == "grep" {
		result, err = grep(t)
	} else {
		result, err = "", fmt.Errorf("command not found %s", c)
	}
	return result, err
}

func cd(s string) (string, error) {
	c := strings.Split(s, " ")
	err := syscall.Chdir(c[1])
	if err != nil {
		return "", fmt.Errorf("%s %s", err, c[1])
	}
	d, err := os.Getwd()
	if err != nil {
		return "", err
	} else {
		return d, nil
	}
}

func ls() (string, error) {
	d, err := os.Getwd()
	if err != nil {
		return "", err
	}
	file, err := os.ReadDir(d)
	if err != nil {
		return "", err
	}
	var b strings.Builder
	for _, f := range file {
		fmt.Fprintf(&b, "%s\n", f.Name())
	}
	return strings.TrimRight(b.String(), "\n"), nil
}

func sl() (string, error) {
	return "∠凸回_曲曲_回回回_回回回~~", nil
}

func cat(s string) (string, error) {
	c := strings.Split(s, " ")
	b, err := os.ReadFile(c[1])
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func pwd() (string, error) {
	d, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return d, nil
}

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

// https://github.com/shirou/gopsutil#usage
func free() (string, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	var b strings.Builder
	_, err = fmt.Fprintf(&b, "Total: %v, Free:%v, UsedPercent:%f%%", v.Total, v.Free, v.UsedPercent)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func echo(t string) (string, error) {
	s := strings.Split(t, " ")
	return s[1], nil
}

func grep(t string) (string, error) {
	s := strings.Split(t, " ")
	if strings.Contains(t, "\n") {
		var b strings.Builder
		for _, ss := range strings.Split(t[len(s[0])+len(s[1])+2:], "\n") {
			if strings.Contains(ss, s[1]) {
				_, err := fmt.Fprintf(&b, "%s\n", ss)
				if err != nil {
					return "", err
				}
			}
		}
		return b.String(), nil
	} else {
		if strings.Contains(s[2], s[1]) {
			return s[2], nil
		} else {
			return "", nil
		}
	}
}
