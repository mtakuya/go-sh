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
	if s[0] == "cd" {
		cd(s)
	} else if s[0] == "ls" {
		ls()
	} else if s[0] == "cat" {
		cat(s)
	} else if s[0] == "pwd" {
		pwd()
	} else {
		fmt.Println("command not found")
	}
}

func exit(err error) {
	fmt.Fprintln(os.Stderr, err)
}

func cd(s []string) {
	err := syscall.Chdir(s[1])
	if err != nil {
		exit(err)
	} else {
		d, err := os.Getwd()
		if err != nil {
			exit(err)
		} else {
			fmt.Println(d)
		}
	}
}

func ls() {
	d, err := os.Getwd()
	if err != nil {
		exit(err)
	}
	file, err := os.ReadDir(d)
	if err != nil {
		exit(err)
	} else {
		for _, f := range file {
			fmt.Println(f.Name())
		}
	}
}

func cat(s []string) {
	b, err := os.ReadFile(s[1])
	if err != nil {
		exit(err)
	} else {
		fmt.Print(string(b))
	}
}

func pwd() {
	d, err := os.Getwd()
	if err != nil {
		exit(err)
	} else {
		fmt.Println(d)
	}
}
