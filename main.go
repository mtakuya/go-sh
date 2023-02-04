package main

import (
	"bufio"
	"fmt"
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
		s.Scan()
		t := s.Text()
		if t == "exit" {
			break
		} else {
			exec(t)
		}
	}
}

func exec(t string) {
	s := strings.Split(t, " ")
	if s[0] == "cd" {
		cd(s)
	} else if s[0] == "ls" {
		ls()
	} else if s[0] == "cat" {
		cat(s)
	} else if s[0] == "pwd" {
		pwd()
	} else {
		os.Exit(1)
	}
}

func exit(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func cd(s []string) {
	syscall.Chdir(s[1])
	d, err := os.Getwd()
	if err != nil {
		exit(err)
	}
	fmt.Println(d)
}

func ls() {
	d, err := os.Getwd()
	if err != nil {
		exit(err)
	}
	file, err := os.ReadDir(d)
	if err != nil {
		exit(err)
	}
	for _, f := range file {
		fmt.Println(f.Name())
	}
}

func cat(s []string) {
	b, err := os.ReadFile(s[1])
	if err != nil {
		exit(err)
	}
	fmt.Print(string(b))
}

func pwd() {
	d, err := os.Getwd()
	if err != nil {
		exit(err)
	}
	fmt.Println(d)
}
