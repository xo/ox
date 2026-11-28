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

- [app get](#app-get): get a file
- [app get all](#app-get-all): get each file
- [app help](#app-help): show help for any command
- [app put](#app-put): put a file
- [app version](#app-version): show app version information

### app get

get a file

#### Synopsis

```
app get [flags] [command] path [dest]
```

#### Options

- `-r, --retries int`: the number of retries (default: 3)
- `--force`: overwrite a file (required)
- `-h, --help`: show help, then exit

#### Arguments

- `path`: the path of the file
- `dest`: where to write the file (default: .)

#### Examples

```
  app get a.txt
```

### app get all

get each file

#### Synopsis

```
app get all [flags] [args]
```

#### Options

- `--dry-run`: write what it gets, and get nothing
- `-h, --help`: show help, then exit

### app help

show help for any command

#### Synopsis

```
app help [flags] [args]
```

#### Description

Help provides help for any app command.

Simply type app help \[path to command\] for full details.

#### Options

- `-h, --help`: show help, then exit

### app put

put a file

#### Synopsis

```
app put [flags] [args]
```

#### Options

- `-h, --help`: show help, then exit

### app version

show app version information

#### Synopsis

```
app version [flags] [args]
```

#### Description

Show app version information.

#### Options

- `-h, --help`: show help, then exit

## Examples

```
  app --config x.yaml get a.txt
  app get -r 5 a.txt /tmp
```

## Notes

See https://example.com/app for more.
