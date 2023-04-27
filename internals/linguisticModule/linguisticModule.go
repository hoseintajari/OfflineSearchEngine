package linguisticModule

import (
	"bufio"
	"regexp"
	"strings"
)

type (
	StopWords struct {
	}
)

func Scanner(s *bufio.Scanner) []string {
	words := bufConvertString(s)
	words = punctuationRemover(words)
	return stopWordsRemover(words)
}

func bufConvertString(s *bufio.Scanner) []string {
	var result []string
	s.Split(bufio.ScanWords)
	for s.Scan() {
		result = append(result, strings.TrimSpace(strings.ToLower(s.Text())))
	}
	return result
}

func punctuationRemover(s []string) []string {
	var result []string
	for _, v := range s {
		cleanWord := regexp.MustCompile(`[^\w]+`).ReplaceAllString(v, "")
		result = append(result, cleanWord)
	}
	return result
}

func stopWordsRemover(s []string) []string {
	var result []string
	stopWords := map[string]bool{
		"a":   true,
		"an":  true,
		"of":  true,
		"the": true,
		"is":  true,
		"and": true,
		"":    true,
	}
	for _, v := range s {
		if !stopWords[v] {
			result = append(result, v)
		}
	}
	return result
}

//old but gold

//func linguisticModule(s *bufio.Scanner) []string {
//	var result []string
//
//	stopWords := map[string]bool{
//		"a":   true,
//		"an":  true,
//		"of":  true,
//		"the": true,
//		"is":  true,
//		"and": true,
//		"":    true,
//	}
//
//	s.Split(bufio.ScanWords)
//	for s.Scan() {
//		word := strings.ToLower(s.Text())
//		word = regexp.MustCompile(`[^\w]+`).ReplaceAllString(word, " ")
//		word = strings.TrimSpace(word)
//		if !stopWords[word] {
//			result = append(result, word)
//		}
//	}
//
//	return result
//}
