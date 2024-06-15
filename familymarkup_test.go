package familymarkup_test

import (
	"context"
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
	syntax "github.com/redexp/tree-sitter-familymarkup"
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
