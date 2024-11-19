package kobra

import (
	"embed"
)

//go:embed bash.sh
//go:embed fish.sh
//go:embed zsh.zsh
//go:embed powershell.ps1
var templates embed.FS
