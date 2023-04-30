package Benchmarks

import (
	"OfflineSearchEngine/internals/SearchEngines/invertedIndex"
	"testing"
)

func BenchmarkSearch(b *testing.B) {
	invertedIndex.BenchmarkInvertedIndexSearch(b)

}
