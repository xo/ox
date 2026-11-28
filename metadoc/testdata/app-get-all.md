# app get all

get each file

## Synopsis

```
app get all [flags] [args]
```

## Options

- `--dry-run`: write what it gets, and get nothing
- `-h, --help`: show help, then exit

## Global Options

- `-r, --retries int`: the number of retries (default: 3)
- `--force`: overwrite a file (required)
- `-c, --config FILE`: the config file (default: $HOME/.apprc)
- `-d, --debug, --trace`: write debug output
- `--format string`: the output format (default: json) (allowed: json, yaml)

## See Also

- [app get](app-get.md): get a file
