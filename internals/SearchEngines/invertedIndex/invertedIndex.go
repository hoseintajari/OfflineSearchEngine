package invertedIndex

//
//import (
//	"OfflineSearchEngine/internals/SearchEngines/models"
//	"bufio"
//)
//
//type InvertedIndex struct {
//	data     map[string]map[int]int
//	docCount int
//}
//
//func NewInvertedIndex() *InvertedIndex {
//	return &InvertedIndex{
//		data:     make(map[string]map[int]int),
//		docCount: 0,
//	}
//}
//
//func (se *InvertedIndexEngine) AddDoc(s *bufio.Scanner, id int) {
//	for Lin {
//		str := s.converter.Convert(s.Text())
//		if str != "" {
//			_, ok := se.data[str]
//			if !ok {
//				se.data[str] = []models.SearchResult{{
//					DocId:         docId,
//					TermFrequency: 1,
//				}}
//			} else {
//				index := se.data[str].Find(docId)
//				if index == -1 {
//					se.data[str] = append(se.data[str], models.SearchResult{
//						DocId:         docId,
//						TermFrequency: 1,
//					})
//				} else {
//					se.data[str][index].TermFrequency++
//				}
//			}
//		}
//	}
//
//}
//
//func (se *InvertedIndexEngine) Search(q string) (models.SearchResults, bool) {
//	res, ok := se.data[q]
//	if !ok {
//		return []models.SearchResult{}, false
//	}
//	return res, true
//}
