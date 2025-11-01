package engine

import (
	_ "embed"
)

//go:embed data/searches.txt
var embeddedDataset []byte

func LoadEmbeddedDataset(e *SuggestionEngine) error {
	return loadFromBytes(embeddedDataset, e)
}
