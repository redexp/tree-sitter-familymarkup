((name) @constant (#is-not? local))

(family
  .
  (name_def
    (name) @constant.builtin
  )
)

(name_ref
  .
  (name) @constant.builtin
  .
  (name)
)

(new_surname
  (name) @constant.builtin
)

(unknown) @keyword

(comment) @comment

(num) @number

[
  ","
  ] @punctuation.delimiter

(sources delimiter: _ @punctuation.delimiter)
(targets delimiter: _ @punctuation.delimiter)

[
  "+"
] @operator

(relation arrow: _ @operator)

(relation label: _ @type)
