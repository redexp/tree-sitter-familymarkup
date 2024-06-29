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
      $.family_name,
      optional(seq(
        $._nl,
        repeatWith($.comment, $._nl)
      )),
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
      choice($.name_ref, $.unknown),
      field('delimiter', choice('+', ',', $._words))
    ),
    targets: $ => repeatWith(
      choice($.name_def, $.unknown, $.num_unknown, $.comment),
      field('delimiter', choice(',', $._nl, $._words))
    ),

    name_ref: $ => seq($.name, optional($.name)),

    family_name: $ => seq(
      $.name,
      optional($.name_aliases)
    ),

    name_def: $ => seq(
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

    comment: _ => seq(choice('*', '/', '#'), optional(/[^\n]+/)),

    unknown: _ => choice('?', /\p{L}[\p{L}\-\d'"\s]*\?/u),
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
