# app

copy files between hosts

## Synopsis

```
app [flags] [command] [args]
```

## Description

App copies files.

It reads its config from FILE.

## Options

- `-c, --config FILE`: the config file (default: $HOME/.apprc)
- `-d, --debug, --trace`: write debug output
- `--format string`: the output format (default: json) (allowed: json, yaml)
- `--token string`: the token
- `-h, --help`: show help, then exit

## Commands

- [app get](app-get.md): get a file
- [app help](app-help.md): show help for any command
- [app put](app-put.md): put a file
- [app version](app-version.md): show app version information

## Examples

```
  app --config x.yaml get a.txt
  app get -r 5 a.txt /tmp
```

## Notes

See https://example.com/app for more.

## See Also

- [app get](app-get.md): get a file
- [app help](app-help.md): show help for any command
- [app put](app-put.md): put a file
- [app version](app-version.md): show app version information
