module github.com/xo/ox/test

go 1.27

require (
	github.com/xo/ox v0.0.0-20260927161223-bb2641a91b18
	github.com/xo/ox/toml v0.0.0-00010101000000-000000000000
	github.com/xo/ox/yaml v0.0.0-00010101000000-000000000000
)

require (
	github.com/goccy/go-yaml v1.19.2 // indirect
	github.com/pelletier/go-toml/v2 v2.4.3 // indirect
)

replace (
	github.com/xo/ox => ../
	github.com/xo/ox/toml => ../toml
	github.com/xo/ox/yaml => ../yaml
)
