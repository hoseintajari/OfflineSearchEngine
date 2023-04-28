package invertedIndex

import (
	"OfflineSearchEngine/internals/SearchEngines/models"
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestCreateInvertedIndex(t *testing.T) {
	ii := CreateInvertedIndex(100)
	if len(ii.Data) != 0 {
		t.Errorf("Expected empty map, got %v", ii.Data)
	}
}
func TestAddDoc(t *testing.T) {
	ii := CreateInvertedIndex(100)
	doc1 := "The quick brown fox jumps over the lazy dog"
	doc2 := "A quick brown dog outpaces a quick brown fox"
	ii.AddDoc(bufio.NewScanner(strings.NewReader(doc1)), 1)
	ii.AddDoc(bufio.NewScanner(strings.NewReader(doc2)), 2)
	results1 := ii.Data["quick"]
	if len(results1) != 2 {
		t.Errorf("Expected 2 search results for 'quick', got %v", results1)
	}
	if results1[0].DocID != 1 || results1[0].Freq != 1 {
		t.Errorf("Expected search result 0 to be {1, 1}, got %v", results1[0])
	}
	if results1[1].DocID != 2 || results1[1].Freq != 2 {
		t.Errorf("Expected search result 1 to be {2, 2}, got %v", results1[1])
	}
}

func TestInvertedIndex_Search(t *testing.T) {
	i := CreateInvertedIndex(10)
	doc1 := "this is a test document"
	doc2 := "A quick brown dog outpaces a quick brown fox test"
	i.AddDoc(bufio.NewScanner(strings.NewReader(doc1)), 1)
	i.AddDoc(bufio.NewScanner(strings.NewReader(doc2)), 2)

	result, ok := i.Search("test")
	if !ok {
		t.Errorf("Expected search result to be true but got false")
	}
	expected := []models.SearchResult{{DocID: 1, Freq: 1}, {DocID: 2, Freq: 1}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected search result to be %v but got %v", expected, result)
	}

	result, ok = i.Search("foobar")
	if ok {
		t.Errorf("Expected search result to be false but got true")
	}
	expected = nil
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected search result to be %v but got %v", expected, result)
	}
}
