package LinearFastAddEngin

import (
	"OfflineSearchEngine/internals/linguisticModule"
	"testing"
)

func BenchmarkLinearFastAddEnginSearch(b *testing.B) {
	var module = linguisticModule.NewLinguisticModule(linguisticModule.ToLower{}, linguisticModule.StopWords{}, linguisticModule.PunctuationRemover{})
	engine := CreateLinearFastAddEngin(10000)
	engine.AddDoc([]string{"foo", "qux", "quux"}, 2, module)
	engine.AddDoc([]string{"quuz", "corge", "foo"}, 3, module)
	engine.AddDoc([]string{"foo", "bar", "baz"}, 4, module)
	engine.AddDoc([]string{"foo", "qux", "quux"}, 5, module)
	engine.AddDoc([]string{"quuz", "corge", "foo"}, 6, module)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = engine.Search("foo")
	}
}
