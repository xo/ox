# app get

get a file

## Synopsis

```
app get [flags] [command] path [dest]
```

## Options

- `-r, --retries int`: the number of retries (default: 3)
- `--force`: overwrite a file (required)
- `-h, --help`: show help, then exit

## Global Options

- `-c, --config FILE`: the config file (default: $HOME/.apprc)
- `-d, --debug, --trace`: write debug output
- `--format string`: the output format (default: json) (allowed: json, yaml)

## Arguments

- `path`: the path of the file
- `dest`: where to write the file (default: .)

## Commands

- [app get all](app-get-all.md): get each file

## Examples

```
  app get a.txt
```

## See Also

- [app](app.md): copy files between hosts
- [app get all](app-get-all.md): get each file
