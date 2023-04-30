package invertedIndex

import (
	"OfflineSearchEngine/internals/linguisticModule"
	"testing"
)

func TestCreateInvertedIndex(t *testing.T) {
	i := CreateInvertedIndex(10)
	if len(i.Data) != 0 {
		t.Errorf("Expected empty data map, but got map with length %d", len(i.Data))
	}
}

func TestAddDoc(t *testing.T) {
	i := CreateInvertedIndex(10)
	s := []string{"sina", "ali", "amin", "test", "document"}
	id := 1
	module := linguisticModule.NewLinguisticModule(linguisticModule.ToLower{}, linguisticModule.StopWords{}, linguisticModule.PunctuationRemover{})
	i.AddDoc(s, id, module)
	if len(i.Data) != 5 {
		t.Errorf("Expected data map with length 5, but got length %d", len(i.Data))
	}
	if len(i.Data["sina"]) != 1 || i.Data["sina"][0].DocID != 1 || i.Data["sina"][0].Freq != 1 {
		t.Errorf("Expected data map to contain {DocID: 1, Freq: 1} for key 'this', but got %v", i.Data["this"])
	}
}

func TestSearch(t *testing.T) {
	i := CreateInvertedIndex(10)
	s1 := []string{"This", "is", "a", "test", "document"}
	id1 := 1
	s2 := []string{"This", "is", "another", "test", "document"}
	id2 := 2
	module := linguisticModule.NewLinguisticModule(linguisticModule.ToLower{}, linguisticModule.StopWords{}, linguisticModule.PunctuationRemover{})
	i.AddDoc(s1, id1, module)
	i.AddDoc(s2, id2, module)
	results, ok := i.Search("test")
	if !ok {
		t.Error("Expected Search to return true, but got false")
	}
	if len(results) != 2 {
		t.Errorf("Expected Search to return 2 results, but got %d", len(results))
	}
	if results[0].DocID != 1 || results[0].Freq != 1 {
		t.Errorf("Expected first result to be {DocID: 1, Freq: 1}, but got %v", results[0])
	}
	if results[1].DocID != 2 || results[1].Freq != 1 {
		t.Errorf("Expected second result to be {DocID: 2, Freq: 1}, but got %v", results[1])
	}
}
