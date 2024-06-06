module.exports = grammar({
	name: 'familymarkup',

	externals: $ => [
		$._nl_eof,
	],

  conflicts: $ => [
    [$.family],
    [$.relations],
  ],

	rules: {
		root: $ => repeatWith($.family, $._multi_newline),

    _multi_newline: _ => /\r?\n\s*\r?\n/,

		family: $ => seq(
      $.name_desc,
      optional(seq(
        $._multi_newline,
        $.relations
      ))
    ),

    relations: $ => repeatWith($.relation, $._multi_newline),

    relation: $ => seq(
      $.sources,
      field('arrow', $._arrows),
      field('label', optional($.words)),
      optional($.targets)
    ),

    sources: $ => repeatWith($.name, field('join', choice('+', ',', $._words))),
    targets: $ => repeatWith($.name_desc, field('join', choice(',', $._words))),

    name_desc: $ => seq(
      $.name,
      optional(
        $.name_aliases
      )
    ),

    name_aliases: $ => seq(
      '(', optional(repeatWith($.name, ',')), ')'
    ),

    name: _ => /\p{Lu}[\p{L}\-\d'"]*/u,
    words: _ => /\p{Ll}[\p{Ll}\s'"]*/u,
    _words: $ => alias($.words, '_words'),

    _arrows: _ => choice('=', '<->', '->', '<-', '-'),
  },
});

function repeatWith(rule, sep) {
  return seq(
    rule,
    repeat(seq(sep, rule))
  );
}
