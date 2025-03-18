//go:build darwin && EMBED_DEPS
// +build darwin,EMBED_DEPS

package lilliput

import "embed"

//go:embed deps/osx/*
//go:embed icc_profiles/*
var EmbedFS embed.FS
