package LinearSortedEngineWithPosting

import (
	"OfflineSearchEngine/internals/SearchEngines/models"
	"OfflineSearchEngine/internals/linguisticModule"
	"sort"
	"strings"
)

type EngineV4 models.EngineV4

func CreateLinearSortedEngineWithPosting(capacity int) *EngineV4 {
	return &EngineV4{Data: make([]models.PostingList, 0, capacity)}
}

func (e *EngineV4) AddDoc(s []string, id int, module linguisticModule.LinguisticModule) {
	wordCount := make(map[string]int)
	for _, word := range module.Convert(s) {
		wordCount[word]++
	}

	for word, freq := range wordCount {
		found := false
		for i, p := range e.Data {
			if p.Term == word {
				e.Data[i].PostingList = append(e.Data[i].PostingList, models.SearchResult{DocID: id, Freq: freq})
				found = true
				break
			}
		}
		if !found {
			e.Data = append(e.Data, models.PostingList{Term: word, PostingList: []models.SearchResult{{DocID: id, Freq: freq}}})
		}
	}

	sort.Slice(e.Data, func(i, j int) bool {
		return e.Data[i].Term < e.Data[j].Term
	})
}

func (e *EngineV4) Search(s string) ([]models.SearchResult, bool) {
	var result []models.SearchResult
	s = strings.ToLower(s)

	for _, v := range e.Data {
		if v.Term == s {
			for _, val := range v.PostingList {
				result = append(result, models.SearchResult{DocID: val.DocID, Freq: val.Freq})
			}
		}
	}

	if len(result) > 0 {
		return result, true
	}

	return result, false
}
