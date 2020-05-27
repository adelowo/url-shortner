package datastore

import (
	"database/sql"
	"io"
)

type postgresDriver struct {
	inner *sql.DB
}

type Store interface {
	io.Closer
}
