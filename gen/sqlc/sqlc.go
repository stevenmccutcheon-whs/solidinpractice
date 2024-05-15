package sqlc

import (
	_ "embed"
)

//go:generate sqlc generate

var (
	//go:embed "sqlite/schemas.sql"
	SQLiteSchemas string
	//go:embed "postgresql/schemas.sql"
	PostgresSchemas string
)
