package app

import "embed"

var (
	//go:embed *.html
	files embed.FS
)
