package tree_sitter_familymarkup_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-familymarkup"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_familymarkup.Language())
	if language == nil {
		t.Errorf("Error loading Familymarkup grammar")
	}
}
