package templates

import "embed"

//go:embed tmpls
var FS embed.FS

//go:embed static
var FSstatic embed.FS
