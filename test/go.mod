module github.com/xo/ox/test

go 1.25.0

require (
	github.com/xo/ox v0.0.0-20261126005800-13bd485f1d7b
	github.com/xo/ox/hcl v0.0.0-00010101000000-000000000000
	github.com/xo/ox/toml v0.0.0-00010101000000-000000000000
	github.com/xo/ox/yaml v0.0.0-00010101000000-000000000000
)

require (
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/apparentlymart/go-textseg/v17 v17.0.1 // indirect
	github.com/goccy/go-yaml v1.19.2 // indirect
	github.com/hashicorp/hcl/v2 v2.25.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/pelletier/go-toml/v2 v2.4.3 // indirect
	github.com/zclconf/go-cty v1.19.0 // indirect
	golang.org/x/mod v0.29.0 // indirect
	golang.org/x/sync v0.18.0 // indirect
	golang.org/x/text v0.31.0 // indirect
	golang.org/x/tools v0.38.0 // indirect
)

replace (
	github.com/xo/ox => ../
	github.com/xo/ox/hcl => ../hcl
	github.com/xo/ox/toml => ../toml
	github.com/xo/ox/yaml => ../yaml
)
