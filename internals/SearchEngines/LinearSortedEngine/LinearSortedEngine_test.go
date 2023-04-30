package LinearSortedEngine

import (
	"OfflineSearchEngine/internals/linguisticModule"
	"testing"
)

func TestCreateLinearSortedEngine(t *testing.T) {
	capacity := 100
	engine := CreateLinearSortedEngine(capacity)
	if len(engine.Data) != 0 || cap(engine.Data) != capacity {
		t.Errorf("Expected engine data to be an empty slice with capacity %d, but got %+v with capacity %d", capacity, engine.Data, cap(engine.Data))
	}
}

var module = linguisticModule.NewLinguisticModule(linguisticModule.ToLower{}, linguisticModule.StopWords{}, linguisticModule.PunctuationRemover{})

func TestAddDoc(t *testing.T) {
	engine := CreateLinearSortedEngine(10)
	if engine == nil {
		t.Error("engine is nil")

		docs := []string{"Hello World", "World is beautiful"}

		engine.AddDoc(docs, 1, module)

		if len(engine.Data) != 5 {
			t.Errorf("expected 5 terms, got %d", len(engine.Data))
		}

		if engine.Data[0].Term != "beautiful" || engine.Data[1].Term != "hello" || engine.Data[2].Term != "is" || engine.Data[3].Term != "world" {
			t.Errorf("expected order of terms to be alphabetical, got %+v", engine.Data)
		}
	}
}
func TestEngineV3_Search(t *testing.T) {
	engine := CreateLinearSortedEngine(10)
	engine.AddDoc([]string{"The", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}, 1, module)
	engine.AddDoc([]string{"A", "brown", "fox", "quickly", "jumps", "over", "the", "lazy", "dog"}, 2, module)
	engine.AddDoc([]string{"The", "lazy", "dog", "is", "always", "lying", "down"}, 3, module)

	result, found := engine.Search("brown")
	if !found {
		t.Errorf("Expected term 'brown' to be found in documents, but not found")
	}
	if len(result) != 2 {
		t.Errorf("Expected to retrieve 2 documents containing term 'fox', but got %d documents", len(result))
	}
	if result[0].DocID != 1 || result[0].Freq != 1 {
		t.Errorf("Expected result[0] to be {DocID: 1, Freq: 1}, but got %v", result[0])
	}
	if result[1].DocID != 2 || result[1].Freq != 1 {
		t.Errorf("Expected result[1] to be {DocID: 2, Freq: 1}, but got %v", result[1])
	}

	result, found = engine.Search("cat")
	if found {
		t.Errorf("Expected term 'cat' not to be found in documents, but found")
	}
	if len(result) != 0 {
		t.Errorf("Expected to retrieve 0 documents containing term 'cat', but got %d documents", len(result))
	}
}
