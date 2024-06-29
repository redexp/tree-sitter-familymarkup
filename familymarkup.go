package familymarkup

import (
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

	count := query.CaptureCount()

	legend := make([]string, count)

	for i := uint32(0); i < count; i++ {
		legend[i] = query.CaptureNameForId(i)
	}

	return legend, nil
}

func GetHighlightCaptures(tree *sitter.Node) ([]*sitter.QueryCapture, error) {
	query, err := GetHighlightQuery()

	if err != nil {
		return nil, err
	}

	cursor := sitter.NewQueryCursor()
	cursor.Exec(query, tree)

	list := make([]*sitter.QueryCapture, 0)
	var prev *sitter.QueryCapture

	for {
		match, ok := cursor.NextMatch()

		if !ok {
			break
		}

		for _, cap := range match.Captures {
			if prev != nil && prev.Node.Equal(cap.Node) {
				if cap.Index > prev.Index {
					list[len(list)-1] = &cap
				}
			} else {
				list = append(list, &cap)
			}

			prev = &cap
		}
	}

	return list, nil
}
