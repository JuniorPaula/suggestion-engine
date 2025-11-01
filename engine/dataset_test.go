package engine

import (
	_ "embed"
)

//go:embed data/test.txt
var embbededTestDataset []byte

func LoadTestDataset(e *SuggestionEngine) error {
	return loadFromBytes(embbededTestDataset, e)
}
