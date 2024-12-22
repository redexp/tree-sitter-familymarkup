package familymarkup

import (
	_ "embed"

	binding "github.com/redexp/tree-sitter-familymarkup/bindings/go"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

//go:embed queries/highlights.scm
var HighlightQuery string

func GetLanguage() *sitter.Language {
	return sitter.NewLanguage(binding.Language())
}

func GetHighlightQuery() (*sitter.Query, *sitter.QueryError) {
	return sitter.NewQuery(GetLanguage(), HighlightQuery)
}

func GetHighlightLegend() ([]string, error) {
	query, err := GetHighlightQuery()

	if err != nil {
		return nil, err
	}

	defer query.Close()

	return query.CaptureNames(), nil
}
