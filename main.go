package main

import (
	"bufio"
	"errors"
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
			} else {
				exec(t)
			}
		}
	}
}

func exec(t string) {
	s := strings.Split(t, " ")
	c := s[0]
	if c == "cd" {
		cd(s)
	} else if c == "ls" {
		ls()
	} else if c == "sl" {
		sl()
	} else if c == "cat" {
		cat(s)
	} else if c == "pwd" {
		pwd()
	} else if c == "ps" {
		ps()
	} else if c == "free" {
		free()
	} else {
		exit(errors.New("command not found"))
	}
}

func exit(err error) {
	fmt.Fprintln(os.Stderr, err)
}

func cd(s []string) {
	err := syscall.Chdir(s[1])
	if err != nil {
		exit(err)
		return
	}
	d, err := os.Getwd()
	if err != nil {
		exit(err)
	} else {
		fmt.Println(d)
	}
}

func ls() {
	d, err := os.Getwd()
	if err != nil {
		exit(err)
		return
	}
	file, err := os.ReadDir(d)
	if err != nil {
		exit(err)
		return
	}
	for _, f := range file {
		fmt.Println(f.Name())
	}
}

func sl() {
	fmt.Println("∠凸回_曲曲_回回回_回回回~~")
}

func cat(s []string) {
	b, err := os.ReadFile(s[1])
	if err != nil {
		exit(err)
		return
	}
	fmt.Print(string(b))
}

func pwd() {
	d, err := os.Getwd()
	if err != nil {
		exit(err)
		return
	}
	fmt.Println(d)
}

func ps() {
	prs, err := process.Processes()
	if err != nil {
		exit(err)
		return
	}
	for _, p := range prs {
		pid := p.Pid
		ppid, err := p.Ppid()
		if err != nil {
			exit(err)
			return
		}
		name, err := p.Name()
		if err != nil {
			exit(err)
			return
		}
		fmt.Println(pid, ppid, name)
	}
}

// https://github.com/shirou/gopsutil#usage
func free() {
	v, err := mem.VirtualMemory()
	if err != nil {
		exit(err)
		return
	}
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
}
