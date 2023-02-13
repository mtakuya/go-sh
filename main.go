package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var histories []string

func main() {
	loop()
}

func loop() {
	var result string
	var err error

	for {
		fmt.Print("> ")
		s := bufio.NewScanner(os.Stdin)

		if s.Scan() {
			t := s.Text()
			histories = append(histories, t)

			if t == "exit" {
				break
			} else if strings.Contains(t, "|") {
				result, err = pipe(t)
			} else {
				result, err = exec(t)
			}

			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			} else {
				if result != "" {
					fmt.Println(result)
				}
			}
		}
	}
}
