package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

const basePath = "/tmp"

func loadFile(name string) (*os.File, error) {
	fpath := filepath.Join(basePath, name)

	f, err := os.Open(fpath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not open file %q", name))
	}

	return f, nil
}

func main() {
	for i := 0; i < 5; i++ {
		f, err := loadFile(strconv.Itoa(i))
		if err != nil {
			fmt.Printf("error: %s\n", err)
			os.Exit(1)
		}

		defer func() {
			if err := f.Close(); err != nil {
				fmt.Printf("error: %s\n", err)
				os.Exit(1)
			}
		}()

		f.WriteString(fmt.Sprintf("Hello in file %d!\n", i))
	}
}
