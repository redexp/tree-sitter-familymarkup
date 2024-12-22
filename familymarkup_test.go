package familymarkup_test

import (
	"testing"

	syntax "github.com/redexp/tree-sitter-familymarkup"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

func TestGetLanguage(t *testing.T) {
	lang := syntax.GetLanguage()
	parser := sitter.NewParser()
	parser.SetLanguage(lang)

	tree := parser.Parse([]byte("Family"), nil)
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

Name1 + Name Surname =
Name2
Name?
`)

	tree := parser.Parse(src, nil)

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

	legend, qerr := syntax.GetHighlightLegend()

	if qerr != nil {
		t.Errorf("GetHighlightLegend: %v", err)
	}

	if legend == nil {
		t.Errorf("legend is nil")
	}

	captures := getHighlightCaptures(tree.RootNode(), query, src)

	compare := [][2]string{
		{"class.declaration.family_name", "Family"},
		{"property.static.name.ref", "Name1"},
		{"operator.sources.join", "+"},
		{"property.static.name.ref", "Name"},
		{"class.family_name.ref", "Surname"},
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

		if cap.Node.Utf8Text(src) != pair[1] {
			t.Errorf("capture %d invalid node %s expect %s", i, cap.Node.Utf8Text(src), pair[1])
		}
	}
}

func getHighlightCaptures(root *sitter.Node, query *sitter.Query, text []byte) []*sitter.QueryCapture {
	cursor := sitter.NewQueryCursor()
	defer cursor.Close()

	list := make([]*sitter.QueryCapture, 0)
	var prev *sitter.QueryCapture

	captures := cursor.Captures(query, root, text)

	for match, idx := captures.Next(); match != nil; match, idx = captures.Next() {
		cap := match.Captures[idx]

		if prev != nil && prev.Node.StartByte() == cap.Node.StartByte() {
			if cap.Index > prev.Index {
				list[len(list)-1] = &cap
			}
		} else {
			list = append(list, &cap)
		}

		prev = &cap
	}

	return list
}
