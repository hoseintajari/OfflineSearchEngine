package CreateEngine

import "OfflineSearchEngine/internals/linguisticModule"

type ISearchEngine struct {
	Engine           SearchEngine
	PathReadDir      string
	DocId            int
	linguisticModule linguisticModule.Converter
}

func CreateISearchEngine(engine SearchEngine, path string, li linguisticModule.Converter) ISearchEngine {
	return ISearchEngine{
		Engine:           engine,
		PathReadDir:      path,
		linguisticModule: li,
		DocId:            0,
	}
}
