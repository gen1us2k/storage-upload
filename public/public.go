package public

import "embed"

// Static stores the latest version of the FE app
//
//go:embed dist/*
var Static embed.FS

// Index stores the latest version of the FE index.html
//
//go:embed dist/index.html
var Index embed.FS
