package LinearFastSearchEngine

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestCreateLinearFastSearchEngine(t *testing.T) {
	// Test case 1: Capacity is zero
	engine := CreateLinearFastSearchEngine(0)
	if len(engine.Data) != 0 {
		t.Errorf("CreateLinearFastSearchEngine: Expected capacity of zero to produce empty data, but got %d", len(engine.Data))
	}

	// Test case 2: Capacity is greater than zero
	engine = CreateLinearFastSearchEngine(10)
	if len(engine.Data) != 0 {
		t.Errorf("CreateLinearFastSearchEngine: Expected capacity of 10 to produce empty data, but got %d", len(engine.Data))
	}
}
func TestAddDoc(t *testing.T) {
	engine := CreateLinearFastSearchEngine(10)
	scanner := bufio.NewScanner(strings.NewReader(""))
	engine.AddDoc(scanner, 1)
	if len(engine.Data) != 0 {
		t.Errorf("AddDoc: Expected empty scanner to produce empty data, but got %d", len(engine.Data))
	}
	engine = CreateLinearFastSearchEngine(10)
	scanner = bufio.NewScanner(strings.NewReader("hello world hello"))
	engine.AddDoc(scanner, 1)
	if len(engine.Data) != 3 {
		t.Errorf("AddDoc: Expected scanner with three words to produce three term info with frequency, but got %d", len(engine.Data))
	}
	if engine.Data[0].DocId != 1 || engine.Data[0].Term != "hello" || engine.Data[0].Freq != 2 {
		t.Errorf("AddDoc: Expected engine data to have term info with DocId=1, Term=hello, and Freq=2, but got %+v", engine.Data[0])
	}
	if engine.Data[1].DocId != 1 || engine.Data[1].Term != "world" || engine.Data[1].Freq != 1 {
		t.Errorf("AddDoc: Expected engine data to have term info with DocId=1, Term=world, and Freq=1, but got %+v", engine.Data[1])
	}
	if engine.Data[2].DocId != 1 || engine.Data[2].Term != "hello" || engine.Data[2].Freq != 2 {
		t.Errorf("AddDoc: Expected engine data to have term info with DocId=1, Term=hello, and Freq=2, but got %+v", engine.Data[2])
	}
}
func TestSearch(t *testing.T) {
	engine := CreateLinearFastSearchEngine(10)

	doc1 := "This is a test document for search engine testing."
	scanner1 := bufio.NewScanner(strings.NewReader(doc1))
	engine.AddDoc(scanner1, 1)

	doc2 := "Search engine testing is important for information retrieval."
	scanner2 := bufio.NewScanner(strings.NewReader(doc2))
	engine.AddDoc(scanner2, 2)

	doc3 := "Search engine testing requires careful consideration of different techniques."
	scanner3 := bufio.NewScanner(strings.NewReader(doc3))
	engine.AddDoc(scanner3, 3)
	results, found := engine.Search("search engine testing")

	if found {
		for _, res := range results {
			fmt.Printf("Document ID: %d, Frequency: %d\n", res.DocID, res.Freq)
		}
	} else {
		fmt.Println("No results found.")
	}
}
