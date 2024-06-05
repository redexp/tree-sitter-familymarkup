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
      $.name,
      '=',
      $.name_desc
    ),

    // source: $ => prec.left(repeatWith($.name, choice('+', ','))),
    // target: $ => prec.left(repeatWith($.name_desc, ',')),

    name_desc: $ => seq(
      $.name,
      optional(
        $.name_aliases
      )
    ),

    name_aliases: $ => seq(
      '(', optional(repeatWith($.name, ',')), ')'
    ),

    name: _ => /\p{Lu}[\p{L}\-\d]*/u,
  },
});

function repeatWith(rule, sep) {
  return seq(
    rule,
    optional(
      seq(sep, repeat(rule))
    )
  );
}
