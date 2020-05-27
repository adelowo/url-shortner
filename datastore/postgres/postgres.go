package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/adelowo/url-shortner/datastore"
)

func New(dsn string) (datastore.Store, error) {

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFn()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return &postgresDriver{inner: db}
}

type postgresDriver struct {
	inner *sql.DB
}

func (p *postgresDriver) Close() error {
	return p.inner.Close()
}
