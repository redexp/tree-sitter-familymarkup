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

    _multi_newline: _ => /\r?\n[\r\n\s]*\r?\n/,
    _nl: _ => /\r?\n/,

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

    sources: $ => repeatWith(
      choice($.name_link, $.unknown),
      choice('+', ',', $._words)
    ),
    targets: $ => repeatWith(
      choice($.name_desc, $.unknown, $.num_unknown),
      choice(',', $._nl, $._words)
    ),

    name_link: $ => seq($.name, optional($.name)),

    name_desc: $ => seq(
      optional($.num),
      optional($.new_surname),
      $.name,
      optional($.name_aliases)
    ),

    num: _ => /\d+[.)]?/,

    new_surname: $ => seq(
      '(', $.name, ')'
    ),

    name_aliases: $ => seq(
      '(', optional(repeatWith($.name, ',')), ')'
    ),

    unknown: _ => /\p{L}?[\p{L}\-\d'"\s]*\?/u,
    num_unknown: $ => seq($.num, $.unknown),

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
