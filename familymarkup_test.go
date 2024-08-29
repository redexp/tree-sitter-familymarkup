package familymarkup_test

import (
	"context"
	"fmt"
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

Name1 + Name? =
Name2
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

	captures, err := syntax.GetHighlightCaptures(tree)

	if err != nil {
		t.Errorf("GetHighlightCaptures: %v", err)
	}

	if captures == nil {
		t.Errorf("captures is nil")
	}

	compare := [][2]string{
		{"constant.builtin.family_name", "Family"},
		{"constant.name.ref", "Name1"},
		{"operator.sources.join", "+"},
		{"keyword.unknown", "Name?"},
		{"operator.arrow", "="},
		{"constant.name.def", "Name2"},
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
