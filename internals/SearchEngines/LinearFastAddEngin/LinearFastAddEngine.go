package LinearFastAddEngin

import (
	"OfflineSearchEngine/internals/SearchEngines/models"
	"OfflineSearchEngine/internals/linguisticModule"
	"strings"
)

type EngineV1 models.EngineV1

func CreateLinearFastAddEngin(capacity int) *EngineV1 {
	return &EngineV1{Data: make([]models.TermInfo, 0, capacity)}
}

func (e *EngineV1) AddDoc(s []string, id int, module linguisticModule.LinguisticModule) {
	for _, v := range module.Convert(s) {
		e.Data = append(e.Data, models.TermInfo{Term: v, DocId: id})
	}
}

func (e *EngineV1) Search(s string) ([]models.SearchResult, bool) {
	s = strings.ToLower(s)
	var result []models.SearchResult

	for i := 0; i < len(e.Data); i++ {
		if s == e.Data[i].Term {
			id := e.Data[i].DocId
			v := find(id, result)
			if v == -1 {
				result = append(result, models.SearchResult{DocID: id, Freq: 1})
			} else {
				result[v].Freq++
			}

		}

	}
	if len(result) > 0 {
		return result, true
	}

	return result, false
}

func find(id int, s []models.SearchResult) int {

	for s, v := range s {
		if v.DocID == id {
			return s
		}
	}
	return -1
}
