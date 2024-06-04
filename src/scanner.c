#include "tree_sitter/parser.h"

enum TokenType {
  NL_OR_EOF
};

void *tree_sitter_familymarkup_external_scanner_create(void) {
  return NULL;
}

void tree_sitter_familymarkup_external_scanner_destroy(void *payload) {
  // ...
}

unsigned tree_sitter_familymarkup_external_scanner_serialize(
  void *payload,
  char *buffer
) {
  return 0;
}

void tree_sitter_familymarkup_external_scanner_deserialize(
  void *payload,
  const char *buffer,
  unsigned length
) {
  // ...
}

bool tree_sitter_familymarkup_external_scanner_scan(
  void *payload,
  TSLexer *lexer,
  const bool *valid_symbols
) {
  if (!valid_symbols[NL_OR_EOF]) return false;

  lexer->result_symbol = NL_OR_EOF;

  while (lexer->lookahead == ' ' || lexer->lookahead == '\t') {
    lexer->advance(lexer, true);
  }

  if (lexer->lookahead == '\n' || lexer->eof(lexer)) {
    return true;
  }

  if (lexer->lookahead == '\r') {
    lexer->advance(lexer, true);
    if (lexer->lookahead == '\n') {
      return true;
    }
  }

  return false;
}
