(sources
  (name_ref (name) @constant.name.ref)
)

(targets
  (name_def (name) @constant.name.ref)
)

(targets
  (name_def 
    (name_aliases (name) @constant.name.def.alias)
  )
)

(relation
  (sources delimiter: "+")
  (targets
    (name_def
      (name) @constant.name.def
    )
  )
)

(family
  .
  (name_def
    (name) @constant.builtin.family_name
  )
)

(family
  .
  (name_def
    (name_aliases
      (name) @constant.builtin.family_name.alias
    )
  )
)

(name_ref
  .
  (name) @constant.builtin.family_name.ref
  .
  (name)
)

(new_surname
  (name) @constant.builtin.family_name.ref
)

(unknown) @keyword.unknown

(comment) @comment

(num) @number.target_num

(sources delimiter: _ @punctuation.delimiter.sources)
(targets delimiter: _ @punctuation.delimiter.targets)

"+" @operator.sources.join

(relation arrow: _ @operator.arrow)

(relation label: _ @type.label)
