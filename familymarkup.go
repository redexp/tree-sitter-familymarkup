package familymarkup

import (
	"bytes"
	_ "embed"

	binding "github.com/redexp/tree-sitter-familymarkup/src"
	sitter "github.com/smacker/go-tree-sitter"
)

//go:embed queries/highlights.scm
var highlightQuery []byte

func GetLanguage() *sitter.Language {
	return sitter.NewLanguage(binding.Language())
}

func GetHighlightQuery() (*sitter.Query, error) {
	return sitter.NewQuery(highlightQuery, GetLanguage())
}

func GetHighlightLegend() ([]string, error) {
	query, err := GetHighlightQuery()

	if err != nil {
		return nil, err
	}

	defer query.Close()

	count := query.CaptureCount()

	legend := make([]string, count)

	for i := uint32(0); i < count; i++ {
		legend[i] = query.CaptureNameForId(i)
	}

	return legend, nil
}
