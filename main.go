package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const comment = "#"
const separator = "="

func main() {}

func Configure(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error while opening env file: %v", err)
	}
	defer func() error {
		if err = file.Close(); err != nil {
			return err
		}
		return nil
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		index := strings.Index(line, comment)
		if index != -1 {
			parsed := strings.Split(line, comment)
			if len(parsed[0]) < 1 {
				continue
			}

			v := strings.Split(parsed[0], separator)
			if len(v[1]) > 1 {
				os.Setenv(v[0], v[1])
			}
			continue
		}
		v := strings.Split(line, separator)
		if len(v[1]) > 1 {
			os.Setenv(v[0], v[1])
		}
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("error while reading env file: %v", err)
	}

	return nil
}
