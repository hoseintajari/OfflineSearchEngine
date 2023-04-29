package fileReader

import (
	"bufio"
	"strings"
)

func bufConvertString(s *bufio.Scanner) []string {
	var result []string
	s.Split(bufio.ScanWords)
	for s.Scan() {
		result = append(result, strings.TrimSpace(strings.ToLower(s.Text())))
	}
	return result
}
