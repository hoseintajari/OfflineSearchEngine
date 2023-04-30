package LinearFastSearchEngine

import (
	"OfflineSearchEngine/internals/SearchEngines/models"
	"OfflineSearchEngine/internals/linguisticModule"
	"strings"
	"testing"
)

func TestCreateLinearFastSearchEngine(t *testing.T) {
	capacity := 10
	engine := CreateLinearFastSearchEngine(capacity)

	if len(engine.Data) != 0 || cap(engine.Data) != capacity {
		t.Errorf("CreateLinearFastSearchEngine did not create an engine with capacity of %d", capacity)
	}
}

var module = linguisticModule.NewLinguisticModule(linguisticModule.ToLower{}, linguisticModule.StopWords{}, linguisticModule.PunctuationRemover{})

func TestAddDoc(t *testing.T) {
	engine := CreateLinearFastSearchEngine(20)
	doc := []string{"ali", "sina", "amin", "test", "document"}
	id := 1
	engine.AddDoc(doc, id, module)

	expectedData := []models.TermInfoWithFreq{
		{Term: "ali", DocId: id, Freq: 1},
		{Term: "sina", DocId: id, Freq: 1},
		{Term: "amin", DocId: id, Freq: 1},
		{Term: "test", DocId: id, Freq: 1},
		{Term: "document", DocId: id, Freq: 1},
	}

	for i, v := range engine.Data {
		if v.Term != expectedData[i].Term || v.DocId != expectedData[i].DocId || v.Freq != expectedData[i].Freq {
			t.Errorf("AddDoc failed. Got %v, expected %v", engine.Data, expectedData)
			return
		}
	}
}

func TestSearch(t *testing.T) {
	engine := CreateLinearFastSearchEngine(100)
	docs := []string{
		"apple pie recipe",
		"how to make apple pie from scratch",
		"easy apple pie recipe",
		"best apple pie recipe",
		"apple pie with crumb topping",
	}
	for i, doc := range docs {
		engine.AddDoc(strings.Split(doc, " "), i, module)
	}

	searchResults, ok := engine.Search("apple pie")
	if !ok {
		t.Errorf("Expected search to return results, got ok=%v", ok)
	}

	expectedResults := []models.SearchResult{
		{DocID: 0, Freq: 1},
		{DocID: 1, Freq: 1},
		{DocID: 2, Freq: 1},
		{DocID: 3, Freq: 1},
		{DocID: 4, Freq: 1},
	}

	if len(searchResults) != len(expectedResults) {
		t.Errorf("Expected %d results, got %d", len(expectedResults), len(searchResults))
	}

	for _, expected := range expectedResults {
		found := false
		for _, result := range searchResults {
			if expected.DocID == result.DocID && expected.Freq == result.Freq {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected result not found: %v", expected)
		}
	}
}

func TestGetFreq(t *testing.T) {
	stringSlice := []string{"this", "is", "a", "test", "document", "this", "is", "is", "test"}
	s := "is"
	freq := GetFreq(s, stringSlice)

	if freq != 3 {
		t.Errorf("GetFreq failed. Got %d, expected %d", freq, 3)
	}
}
