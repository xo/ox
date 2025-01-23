package ox

import "embed"

// templates are the embedded completion templates.
//
//go:embed comp/*
var templates embed.FS
