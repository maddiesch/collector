package stubbed

import "embed"

//go:embed data/*
var DataFS embed.FS
