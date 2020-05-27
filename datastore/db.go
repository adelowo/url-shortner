package datastore

import (
	"context"
	"database/sql"
	"io"
	"net/url"
	"time"
)

type postgresDriver struct {
	inner *sql.DB
}

type Store interface {
	io.Closer
	Create(context.Context, URL) error
	Find(ctx context.Context, code string) (*URL, error)
}

type URL struct {
	Code        string
	RedirectURL *url.URL
	CreatedAt   time.Time
}
