package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
