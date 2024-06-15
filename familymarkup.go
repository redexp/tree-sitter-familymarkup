package familymarkup

import (
	sitter "github.com/smacker/go-tree-sitter"
	binding "github.com/redexp/tree-sitter-familymarkup/src"
)

func GetLanguage() *sitter.Language {
	return sitter.NewLanguage(binding.Language())
}
