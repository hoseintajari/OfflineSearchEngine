package linguisticModule

import (
	"strings"
)

type Converter interface {
	Convert(s []string) []string
}

type (
	StopWords struct {
	}
	PunctuationRemover struct {
	}
	ToLower struct {
	}
)

func (st StopWords) Convert(s []string) []string {
	var result []string
	stopWords := map[string]bool{
		"a":   true,
		"an":  true,
		"of":  true,
		"the": true,
		"is":  true,
		"and": true,
		"to":  true,
	}
	for _, v := range s {
		if !stopWords[v] {
			result = append(result, v)
		}
	}
	return result
}

func (st PunctuationRemover) Convert(s []string) []string {
	var result []string
	punctuations := map[string]bool{
		".": true,
		",": true,
		";": true,
		":": true,
		"!": true,
		"?": true,
		`"`: true,
		`'`: true,
		")": true,
		"]": true,
		"}": true,
		">": true,
		"(": true,
		"[": true,
		"{": true,
		"<": true,
	}

	for _, v := range s {
		if !punctuations[string(v[len(v)-1])] {
			result = append(result, v)
		} else {
			result = append(result, v[:len(v)-1])
		}
	}
	return result
}

func (st ToLower) Convert(s []string) []string {
	var result []string
	for _, v := range s {
		result = append(result, strings.ToLower(v))
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
