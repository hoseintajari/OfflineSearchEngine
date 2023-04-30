package LinearFastSearchEngine

import (
	"OfflineSearchEngine/internals/linguisticModule"
	"testing"
)

func BenchmarkSearch(b *testing.B) {
	module := linguisticModule.NewLinguisticModule(linguisticModule.ToLower{}, linguisticModule.StopWords{}, linguisticModule.PunctuationRemover{})

	engine := CreateLinearFastSearchEngine(1000)
	engine.AddDoc([]string{"foo", "qux", "quux"}, 2, module)
	engine.AddDoc([]string{"quuz", "corge", "foo"}, 3, module)
	engine.AddDoc([]string{"foo", "bar", "baz"}, 4, module)
	engine.AddDoc([]string{"foo", "qux", "quux"}, 5, module)
	engine.AddDoc([]string{"quuz", "corge", "foo"}, 6, module)
	searchText := "test"
	for n := 0; n < b.N; n++ {
		_, _ = engine.Search(searchText)
	}
}
