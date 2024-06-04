module.exports = grammar({
	name: 'familymarkup',

	externals: $ => [
		$._nl_eof,
	],

  conflicts: $ => [[$.family]],

	rules: {
		root: $ => repeat($.family),

		family: $ => seq(
      $.name,
      $._nl_eof,
      repeat($.relation)
    ),

    relation: $ => seq(
      $.name, '=', $.name, $._nl_eof
    ),

    name: _ => /\p{Lu}[\p{L}\-\d]*/u,
  },
});
