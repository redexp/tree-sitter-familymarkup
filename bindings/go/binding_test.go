package tree_sitter_familymarkup_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_familymarkup "github.com/redexp/tree-sitter-familymarkup/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_familymarkup.Language())
	if language == nil {
		t.Errorf("Error loading familymarkup grammar")
	}
}
