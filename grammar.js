module.exports = grammar({
	name: 'familymarkup',

  conflicts: $ => [
    [$.family],
    [$.relations],
    [$.relation],
    [$.targets],
  ],

  extras: _ => [' ', '\t'],

	rules: {
		root: $ => seq(
      optional(choice($._nl, $._multi_newline)),
      repeatWith($.family, $._multi_newline),
      optional(choice($._nl, $._multi_newline)),
    ),

    _multi_newline: _ => prec(2, /\r?\n[\r\n\s]*\r?\n/),
    _nl: _ => prec(1, /\r?\n/),

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
      optional($._nl),
      optional($.targets),
    ),

    sources: $ => repeatWith($.name_link, field('join', choice('+', ',', $._words))),
    targets: $ => repeatWith($.name_desc, field('join', choice(',', $._nl, $._words))),

    name_link: $ => seq($.name, optional($.name)),

    name_desc: $ => seq(
      optional($.new_surname),
      $.name,
      optional($.name_aliases)
    ),

    new_surname: $ => seq(
      '(', $.name, ')'
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
    optional(repeat(seq(sep, rule)))
  );
}
