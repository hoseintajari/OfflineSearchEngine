package LinearSortedEngine

import (
	"OfflineSearchEngine/internals/SearchEngines/models"
	"OfflineSearchEngine/internals/linguisticModule"
	"bufio"
	"sort"
	"strings"
)

type EngineV3 models.EngineV2

func CreateLinearSortedEngine(capacity int) *EngineV3 {
	return &EngineV3{Data: make([]models.TermInfoWithFreq, 0, capacity)}
}

func (e EngineV3) AddDoc(s *bufio.Scanner, id int) {
	stringSlice := linguisticModule.Scanner(s)
	for _, v := range stringSlice {
		e.Data = append(e.Data, models.TermInfoWithFreq{Term: v, DocId: id, Freq: GetFreq(v, stringSlice)})
	}
	sort.SliceStable(e.Data, func(i, j int) bool {
		if e.Data[i].Term != e.Data[j].Term {
			return e.Data[i].Term < e.Data[j].Term
		}
		return e.Data[i].Term < e.Data[j].Term
	})
}

func (e EngineV3) Search(s string) ([]models.SearchResult, bool) {
	var result []models.SearchResult
	s = strings.ToLower(s)
	for _, v := range e.Data {
		condition := find(v.DocId, result)
		if condition != true {
			result = append(result, models.SearchResult{DocID: v.DocId, Freq: v.Freq})
		}
	}
	if len(result) > 0 {
		return result, true
	}
	return result, false
}

func find(id int, s []models.SearchResult) bool {

	for _, v := range s {
		if v.DocID == id {
			return true
		}
	}
	return false
}

func GetFreq(s string, stringSLice []string) int {
	Frecuency := 0
	for _, v := range stringSLice {
		if s == v {
			Frecuency++
		}
	}
	return Frecuency
}
