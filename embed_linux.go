//go:build linux && EMBED_DEPS
// +build linux,EMBED_DEPS

package lilliput

import "embed"

//go:embed deps/linux/*
//go:embed icc_profiles/*
var EmbedFS embed.FS
