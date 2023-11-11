package app

import "embed"

var (
	//go:embed *.html
	templateFiles embed.FS
)

var (
	//go:embed  static/*
	staticFiles embed.FS
)
