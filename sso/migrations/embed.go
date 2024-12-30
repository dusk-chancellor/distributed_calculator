package migrations

import "embed"

type FS struct {
	FS embed.FS
	Dir string
}
//go:embed sql/*.sql
var fs embed.FS

var Migrations = &FS{
	FS: fs,
	Dir: "sql",
}
