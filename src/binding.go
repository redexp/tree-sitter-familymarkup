package src

// #cgo CFLAGS: -std=c11 -fPIC
// #include "./tree_sitter/parser.h"
//TSLanguage *tree_sitter_familymarkup();
import "C"

import "unsafe"

// Get the tree-sitter Language for this grammar.
func Language() unsafe.Pointer {
	return unsafe.Pointer(C.tree_sitter_familymarkup())
}
