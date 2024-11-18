package familymarkup_test

import (
	"context"
	"testing"

	syntax "github.com/redexp/tree-sitter-familymarkup"
	sitter "github.com/smacker/go-tree-sitter"
)

func TestGetLanguage(t *testing.T) {
	lang := syntax.GetLanguage()
	parser := sitter.NewParser()
	parser.SetLanguage(lang)

	input := []byte("Family")
	tree, err := sitter.ParseCtx(context.Background(), input, lang)
	if err != nil {
		t.Errorf("error parsing input: %v", err)
	}
	if tree == nil {
		t.Errorf("tree is nil")
	}
}

func TestGetHighlightQuery(t *testing.T) {
	lang := syntax.GetLanguage()
	parser := sitter.NewParser()
	parser.SetLanguage(lang)

	src := []byte(`
Family

Name1 + Surname Name =
Name2
Name?
`)

	tree, _ := sitter.ParseCtx(context.Background(), src, lang)

	if tree == nil {
		t.Errorf("tree is nil")
	}

	query, err := syntax.GetHighlightQuery()

	if err != nil {
		t.Errorf("GetHighlightQuery: %v", err)
	}

	if query == nil {
		t.Errorf("query is nil")
	}

	legend, err := syntax.GetHighlightLegend()

	if err != nil {
		t.Errorf("GetHighlightLegend: %v", err)
	}

	if legend == nil {
		t.Errorf("legend is nil")
	}

	captures := getHighlightCaptures(tree, query)

	if captures == nil {
		t.Errorf("captures is nil")
	}

	compare := [][2]string{
		{"class.declaration.family_name", "Family"},
		{"property.static.name.ref", "Name1"},
		{"operator.sources.join", "+"},
		{"class.family_name.ref", "Surname"},
		{"property.static.name.ref", "Name"},
		{"operator.arrow", "="},
		{"property.declaration.static.name.def", "Name2"},
		{"string.unknown", "Name?"},
	}

	if len(captures) != len(compare) {
		t.Errorf("captures invalid length")
	}

	for i, pair := range compare {
		cap := captures[i]

		if legend[cap.Index] != pair[0] {
			t.Errorf("capture %d invalid type %s should be %s", i, legend[captures[i].Index], pair[0])
		}

		if cap.Node.Content(src) != pair[1] {
			t.Errorf("capture %d invalid node %s expect %s", i, cap.Node.Content(src), pair[1])
		}
	}
}

func getHighlightCaptures(tree *sitter.Node, query *sitter.Query) []*sitter.QueryCapture {
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

	return list
}
