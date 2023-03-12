package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func md5sum(s string) (string, error) {
	c := strings.Split(s, " ")
	if len(c) != 2 {
		return "", errors.New("md5sum error")
	}

	if _, err := os.Stat(c[1]); err == nil {
		f, err := os.Open(c[1])
		if err != nil {
			return "", err
		}
		defer func() {
			err := f.Close()
			if err != nil {
				fmt.Println(err)
			}
		}()

		h := md5.New()
		if _, err := io.Copy(h, f); err != nil {
			return "", err
		}
		return hex.EncodeToString(h.Sum(nil)), nil
	} else if errors.Is(err, os.ErrNotExist) {
		h := md5.New()
		_, err := io.WriteString(h, c[1])
		if err != nil {
			return "", err
		}
		return hex.EncodeToString(h.Sum(nil)), nil
	} else {
		return "", err
	}
}
