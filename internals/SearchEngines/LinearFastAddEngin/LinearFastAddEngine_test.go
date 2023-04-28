package LinearFastAddEngin

import (
	"bufio"
	"strings"
	"testing"
)

func TestCreateLinearFastAddEngin(t *testing.T) {
	capacity := 10
	engine := CreateLinearFastAddEngin(capacity)

	if len(engine.Data) != 0 {
		t.Errorf("Expected an empty data slice but got %v instead", engine.Data)
	}

	if cap(engine.Data) != capacity {
		t.Errorf("Expected capacity of %v but got %v instead", capacity, cap(engine.Data))
	}
}

//func TestEngineV1_AddDoc(t *testing.T) {
//
//	engine := CreateLinearFastAddEngin(10)
//	doc1 := "Hello world, welcome to my world."
//	doc2 := "The world is a beautiful place to live in."
//	doc3 := "This is a test document."
//	engine.AddDoc(bufio.NewScanner(strings.NewReader(doc1)), 1)
//	engine.AddDoc(bufio.NewScanner(strings.NewReader(doc2)), 2)
//	engine.AddDoc(bufio.NewScanner(strings.NewReader(doc3)), 3)
//
//	expectedData := []models.TermInfo{
//		{Term: "hello", DocId: 1},
//		{Term: "world", DocId: 1},
//		{Term: "welcome", DocId: 1},
//		{Term: "world", DocId: 1},
//		{Term: "the", DocId: 2},
//		{Term: "world", DocId: 2},
//		{Term: "is", DocId: 2},
//		{Term: "a", DocId: 2},
//		{Term: "beautiful", DocId: 2},
//		{Term: "place", DocId: 2},
//		{Term: "live", DocId: 2},
//		{Term: "in", DocId: 2},
//		{Term: "this", DocId: 3},
//		{Term: "is", DocId: 3},
//		{Term: "a", DocId: 3},
//		{Term: "test", DocId: 3},
//		{Term: "document", DocId: 3},
//	}
//	if len(engine.Data) != len(expectedData) {
//		t.Errorf("Expected %d terms, but got %d", len(expectedData), len(engine.Data))
//	}
//
//	for i := range engine.Data {
//		if engine.Data[i] != expectedData[i] {
//			t.Errorf("Expected %+v, but got %+v", expectedData[i], engine.Data[i])
//		}
//	}
//}

func TestEngineV1_Search(t *testing.T) {
	engine := CreateLinearFastAddEngin(10)

	doc1 := bufio.NewScanner(strings.NewReader("this is the first document"))
	engine.AddDoc(doc1, 1)
	doc2 := bufio.NewScanner(strings.NewReader("this is the second document"))
	engine.AddDoc(doc2, 2)

	results, found := engine.Search("this")
	if !found {
		t.Errorf("expected true, got %t", found)
	}

	if len(results) != 2 {
		t.Errorf("expected 2 results, got %d", len(results))
	}

	if results[0].DocID != 1 {
		t.Errorf("expected DocID 1, got %d", results[0].DocID)
	}

	if results[0].Freq != 1 {
		t.Errorf("expected Freq 1, got %d", results[0].Freq)
	}

	if results[1].DocID != 2 {
		t.Errorf("expected DocID 2, got %d", results[1].DocID)
	}

	if results[1].Freq != 1 {
		t.Errorf("expected Freq 1, got %d", results[1].Freq)
	}
	results, found = engine.Search("third")
	if found {
		t.Errorf("expected false, got %t", found)
	}
	if len(results) != 0 {
		t.Errorf("expected 0 results, got %d", len(results))
	}
}
