# Notes

## Diff detection

As we perform the generation - the runtime is not idempotent - as time will change after each invocation (same for UUID)

Can we generate the document and then perform a diff if there is an existing document present?

## Hashes

Add functionality after aggregation is stable to allow the hash identification of the file in the declaration.
IE:

```yaml
name: my-generated-file.yaml
components:
    local:
    - name: component-1.yaml
      hash: <known hash>
```

## TODO

- Add ability to override version on the Command Line
- Proper error checking everywhere
- consistent logging
