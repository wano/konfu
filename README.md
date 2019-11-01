# konfu

## what is this？

pongo2 template render command.

## how to install

```bash
go get -u github/wano/konfu
```

## usage

```bash
konfu -c test/sample.json -t test/sample.tpl -o /path/to/output
```

- -c : json parameter file path(required)
- -t : template file path(required)
- -o : output file path(if not specified, output to stdout)