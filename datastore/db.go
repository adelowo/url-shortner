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
	Create(context.Context, *URL) (string, error)
	Find(ctx context.Context, code string) (*URL, error)
}

type URL struct {
	ID          int64     `json:"id"`
	Code        string    `json:"code"`
	RedirectURL *url.URL  `json:"redirect_url"`
	CreatedAt   time.Time `json:"created_at"`
}
