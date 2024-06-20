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

func GetHighlightLegend() (map[uint32]string, error) {
	query, err := GetHighlightQuery()

	if err != nil {
		return nil, err
	}

	legend := make(map[uint32]string)

	count := query.CaptureCount()
	for i := uint32(0); i < count; i++ {
		legend[i] = query.CaptureNameForId(i)
	}

	return legend, nil
}

func GetHighlightCaptures(tree *sitter.Node) (list []sitter.QueryCapture, err error) {
	query, err := GetHighlightQuery()

	if err != nil {
		return nil, err
	}

	cursor := sitter.NewQueryCursor()
	cursor.Exec(query, tree)

	var prev *sitter.QueryCapture

	for {
		match, ok := cursor.NextMatch()

		if !ok {
			break
		}

		for _, cap := range match.Captures {
			if prev != nil && prev.Node.Equal(cap.Node) {
				if cap.Index > prev.Index {
					list[len(list)-1] = cap
				}
			} else {
				list = append(list, cap)
			}

			prev = &cap
		}
	}

	return list, nil
}
