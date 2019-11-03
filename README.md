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

```
Usage of konfu:
  -c string
    	parameter file path
  -m string
    	parameter file mode. json or yaml (default "json")
  -o string
    	render result output file path. if not specified, output to stdout.
  -t string
    	template file path
```