package LinearFastAddEngin

import (
	"OfflineSearchEngine/internals/linguisticModule"
	"testing"
)

var module = linguisticModule.NewLinguisticModule(linguisticModule.ToLower{}, linguisticModule.StopWords{}, linguisticModule.PunctuationRemover{})

func TestCreateLinearFastAddEngin(t *testing.T) {
	capacity := 10
	engine := CreateLinearFastAddEngin(capacity)
	if len(engine.Data) != 0 || cap(engine.Data) != capacity {
		t.Errorf("Unexpected capacity or length for EngineV1")
	}
}

func TestAddDoc(t *testing.T) {
	engine := CreateLinearFastAddEngin(10)
	engine.AddDoc([]string{"hello", "world", "hello"}, 1, module)
	engine.AddDoc([]string{"world", "world", "world"}, 2, module)
	engine.AddDoc([]string{"foo", "bar", "baz"}, 3, module)

	if len(engine.Data) != 9 {
		t.Errorf("Expected 9 terms in EngineV1 data, but got %d", len(engine.Data))
	}
}

func TestSearch(t *testing.T) {

	engine := CreateLinearFastAddEngin(10)
	engine.AddDoc([]string{"hello", "world", "hello"}, 1, module)
	engine.AddDoc([]string{"world", "world", "world"}, 2, module)
	engine.AddDoc([]string{"foo", "bar", "baz"}, 3, module)

	// Test a search for a term that exists in multiple documents
	results, ok := engine.Search("world")
	if !ok {
		t.Errorf("Expected results for search term 'world'")
	}
	if len(results) != 2 {
		t.Errorf("Expected 2 search results for term 'world', but got %d", len(results))
	}
	if results[0].DocID != 1 || results[0].Freq != 1 {
		t.Errorf("Unexpected search result for DocID and Freq for term 'world'")
	}
	if results[1].DocID != 2 || results[1].Freq != 3 {
		t.Errorf("Unexpected search result for DocID and Freq for term 'world'")
	}

	// Test a search for a term that does not exist in any document
	results, ok = engine.Search("foobar")
	if ok {
		t.Errorf("Expected no results for search term 'foobar'")
	}
	if len(results) != 0 {
		t.Errorf("Expected 0 search results for term 'foobar', but got %d", len(results))
	}
}
