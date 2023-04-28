package LinearFastSearchEngine

import (
	"OfflineSearchEngine/internals/SearchEngines/models"
	"OfflineSearchEngine/internals/linguisticModule"
	"bufio"
	"strings"
)

type EngineV2 models.EngineV2

func CreateLinearFastSearchEngine(capacity int) *EngineV2 {
	return &EngineV2{Data: make([]models.TermInfoWithFreq, 0, capacity)}
}

func (e *EngineV2) AddDoc(s *bufio.Scanner, id int) {
	stringSlice := linguisticModule.Scanner(s)
	for _, v := range stringSlice {
		e.Data = append(e.Data, models.TermInfoWithFreq{Term: v, DocId: id, Freq: GetFreq(v, stringSlice)})
	}
}

func (e *EngineV2) Search(s string) ([]models.SearchResult, bool) {
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
	frequency := 0
	for _, v := range stringSLice {
		if s == v {
			frequency++
		}
	}
	return frequency
}
