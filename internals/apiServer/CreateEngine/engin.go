package CreateEngine

import (
	"OfflineSearchEngine/internals/SearchEngines/LinearFastAddEngin"
	"OfflineSearchEngine/internals/SearchEngines/LinearFastSearchEngine"
	"OfflineSearchEngine/internals/SearchEngines/LinearSortedEngine"
	"OfflineSearchEngine/internals/SearchEngines/LinearSortedEngineWithPosting"
	"OfflineSearchEngine/internals/SearchEngines/models"
	"bufio"
)

type SearchEngine interface {
	AddDoc(s *bufio.Scanner, id int)
	Search(s string) ([]models.SearchResult, bool)
}

func NewSearchEngine(str string, cap int) SearchEngine {
	switch str {
	case "v1":
		return LinearFastAddEngin.CreateLinearFastAddEngin(cap)
	case "v2":
		return LinearFastSearchEngine.CreateLinearFastSearchEngine(cap)
	case "v3":
		return LinearSortedEngine.CreateLinearSortedEngine(cap)
	case "v4":
		return LinearSortedEngineWithPosting.CreateLinearSortedEngineWithPosting(cap)
	default:
		return nil
	}
}
