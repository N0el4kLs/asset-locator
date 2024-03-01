package util

import (
	"bufio"
	"os"
	"strings"
)

func ReadFileByLine(location string) ([]string, error) {
	var lines []string
	file, err := os.Open(location)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		lines = append(lines, line)
	}
	return lines, nil
}
