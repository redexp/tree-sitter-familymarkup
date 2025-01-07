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
      field('name', $.family_name),
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
      field('sources', $.sources),
      field('arrow', choice($.eq, $.arrow)),
      field('label', optional($.words)),
      optional($._nl),
      field('targets', optional($.targets)),
    ),

    sources: $ => repeatWith(
      choice($.name_ref, $.name, $.unknown),
      field('delimiter', choice('+', ',', $._words))
    ),
    targets: $ => repeatWith(
      choice($.name_def, $.num_unknown, $.unknown, $.comment),
      field('delimiter', choice(',', $._nl, $._words))
    ),

    name_ref: $ => seq($.name, alias($.name, $.surname)),

    family_name: $ => seq(
      field('name', $.name),
      field('aliases', optional($.name_aliases)),
    ),

    name_def: $ => seq(
      field('number', optional($.num)),
      field('name', $.name),
      field('aliases', optional($.name_aliases)),
      field('surname', optional(alias($.name, $.surname))),
    ),

    num_unknown: $ => seq(
      field('number', $.num),
      field('name', $.unknown),
    ),

    num: _ => /\d+[.)]?/,

    name_aliases: $ => seq(
      '(', optional(repeatWith($.name, ',')), ')'
    ),

    comment: _ => seq(choice('*', '/', '#'), optional(/[^\n]+/)),

    name: _ => /\p{Lu}[\p{L}\-\d'".]*/u,

    unknown: _ => choice('?', /\p{L}[\p{L}\-\d'" ]*\?/u),

    words: _ => /\p{Ll}([\p{Ll}'"\s]*[\p{Ll}'"])?/u,
    _words: $ => alias($.words, '_words'),

    eq: _ => "=",

    arrow: _ => choice('<->', '->', '<-', '-'),
  },
});

function repeatWith(rule, sep) {
  return seq(
    rule,
    optional(repeat(seq(sep, rule)))
  );
}
