package invertedIndex

import (
	"OfflineSearchEngine/internals/SearchEngines/models"
	"OfflineSearchEngine/internals/linguisticModule"
)

type InvertedIndex models.InvertedIndex

func CreateInvertedIndex(capacity int) *InvertedIndex {
	return &InvertedIndex{Data: make(map[string][]models.SearchResult, capacity)}
}

func (i *InvertedIndex) AddDoc(s []string, id int, module linguisticModule.LinguisticModule) {
	for _, word := range module.Convert(s) {
		_, ok := i.Data[word]
		if !ok {
			i.Data[word] = []models.SearchResult{{DocID: id, Freq: 1}}
		} else {
			last := len(i.Data[word]) - 1
			if i.Data[word][last].DocID == id {
				i.Data[word][last].Freq++
			} else {
				i.Data[word] = append(i.Data[word], models.SearchResult{DocID: id, Freq: 1})
			}
		}
	}
}

func (i *InvertedIndex) Search(s string) ([]models.SearchResult, bool) {
	result, ok := i.Data[s]
	if !ok {
		return nil, false
	}
	return result, true
}
