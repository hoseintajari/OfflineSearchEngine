package LinearSortedEngineWithPosting

import (
	"OfflineSearchEngine/internals/SearchEngines/models"
	"OfflineSearchEngine/internals/linguisticModule"
	"reflect"
	"testing"
)

var module = linguisticModule.NewLinguisticModule(linguisticModule.ToLower{}, linguisticModule.StopWords{}, linguisticModule.PunctuationRemover{})

func TestCreateLinearSortedEngineWithPosting(t *testing.T) {
	capacity := 100
	engine := CreateLinearSortedEngineWithPosting(capacity)

	if len(engine.Data) != 0 {
		t.Errorf("CreateLinearSortedEngineWithPosting failed: expected %d, got %d", 0, len(engine.Data))
	}

	if cap(engine.Data) != capacity {
		t.Errorf("CreateLinearSortedEngineWithPosting failed: expected %d, got %d", capacity, cap(engine.Data))
	}
}

func TestAddDoc(t *testing.T) {
	engine := CreateLinearSortedEngineWithPosting(10)

	doc1 := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	doc2 := []string{"the", "lazy", "dog", "jumps", "over", "the", "quick", "brown", "fox"}

	engine.AddDoc(doc1, 1, module)
	engine.AddDoc(doc2, 2, module)

	expectedData := []models.PostingList{
		{Term: "brown", PostingList: []models.SearchResult{{DocID: 1, Freq: 1}, {DocID: 2, Freq: 1}}},
		{Term: "dog", PostingList: []models.SearchResult{{DocID: 1, Freq: 1}, {DocID: 2, Freq: 1}}},
		{Term: "fox", PostingList: []models.SearchResult{{DocID: 1, Freq: 1}, {DocID: 2, Freq: 1}}},
		{Term: "jumps", PostingList: []models.SearchResult{{DocID: 1, Freq: 1}, {DocID: 2, Freq: 1}}},
		{Term: "lazy", PostingList: []models.SearchResult{{DocID: 1, Freq: 1}, {DocID: 2, Freq: 1}}},
		{Term: "over", PostingList: []models.SearchResult{{DocID: 1, Freq: 1}, {DocID: 2, Freq: 1}}},
		{Term: "quick", PostingList: []models.SearchResult{{DocID: 1, Freq: 1}, {DocID: 2, Freq: 1}}},
	}

	if !reflect.DeepEqual(engine.Data, expectedData) {
		t.Errorf("AddDoc failed: expected %v, got %v", expectedData, engine.Data)
	}
}
func TestSearch(t *testing.T) {
	e := CreateLinearSortedEngineWithPosting(100)

	doc1 := []string{"Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit"}
	doc2 := []string{"Sed", "ut", "perspiciatis", "unde", "omnis", "iste", "natus", "error"}

	e.AddDoc(doc1, 1, module)
	e.AddDoc(doc2, 2, module)

	result, ok := e.Search("lorem")
	if !ok {
		t.Errorf("Expected ok to be true, but got false")
	}
	expected := []models.SearchResult{{DocID: 1, Freq: 1}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	result, ok = e.Search("natus")
	if !ok {
		t.Errorf("Expected ok to be true, but got false")
	}
	expected = []models.SearchResult{{DocID: 2, Freq: 1}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	result, ok = e.Search("nonexistent")
	if ok {
		t.Errorf("Expected ok to be false, but got true")
	}
	expected = []models.SearchResult{}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
