# helm-unused-values

Helm plugin to find unused values in `values.yaml`

Different kinds of unused values in `values.yaml`

- The default `values.yaml` has values that are not being used at all in the
  templates. This maybe because of various reasons. Two examples:
  - The value is an old value is not used anymore. More like deprecated now, or
    just dead code.
  - Some typo issue or some naming issue. Ideally, the templating operation
    should fail in this case, saying that the value it was looking for was not
    found - because it's named differently in the default `values.yaml`
- The overriding `values.yaml` has values that are not being used at all in
  the templates. Again, this could also be because of various reasons. Assume
  similar examples like the above

How to solve these? Some of these issues can be solved by `values.schema.json`,
by using some strict validations to avoid unknown properties and helm's
templating engine which can error out when the values it needs are not found.
Maybe strict templating mode can be used for more vigorous checks

What if the `values.schema.json` has the validation for the value and the value
is also present in default `values.yaml` and maybe also in the external
`values.yaml` but is not being used by the chart? That's dead code :)
