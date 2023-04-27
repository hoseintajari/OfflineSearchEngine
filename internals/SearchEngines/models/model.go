package models

type (
	SearchResult struct {
		DocID int
		Freq  int
	}
	TermInfo struct {
		Term  string
		DocId int
	}
	PostingList struct {
		Term        string
		PostingList []SearchResult
	}
	TermInfoWithFreq struct {
		Term  string
		DocId int
		Freq  int
	}
)
type (
	EngineV1 struct {
		Data []TermInfo
	}
	EngineV2 struct {
		Data []TermInfoWithFreq
	}
	EngineV4 struct {
		Data []PostingList
	}
	InvertedIndex struct {
		Data map[string][]SearchResult
	}
)
