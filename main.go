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
			if t == "exit" {
				break
			} else if strings.Contains(t, "|") {
				histories = append(histories, t)
				result, err = pipe(t)
			} else if 0 < len(t) && t[0] == '!' {
				result, err = historyExec(t)
			} else {
				histories = append(histories, t)
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
