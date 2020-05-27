package postgres

import (
	"context"
	"database/sql"
	"net/url"
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

	return &postgresDriver{inner: db}, nil
}

type postgresDriver struct {
	inner *sql.DB
}

func (p *postgresDriver) Close() error {
	return p.inner.Close()
}

func (p *postgresDriver) Find(ctx context.Context, redirectURL string) (*datastore.URL, error) {
	var u = new(datastore.URL)

	row := p.inner.QueryRowContext(ctx, "SELECT id,redirect_url,code FROM url WHERE redirect_url = $1 ", redirectURL)

	if err := row.Scan(&u.ID, &redirectURL, &u.Code); err != nil {
		return nil, err
	}

	u.RedirectURL, _ = url.Parse(redirectURL)
	return u, nil
}

func (p *postgresDriver) Create(ctx context.Context, val *datastore.URL) (string, error) {
	u, err := p.Find(ctx, val.RedirectURL.String())
	if err != nil {
		if err == sql.ErrNoRows {
			_, err = p.inner.ExecContext(ctx,
				"INSERT INTO url(code,redirect_url,created_at) VALUES($1,$2,$3)",
				val.Code, val.RedirectURL.String(), val.CreatedAt)
			return "", err
		}

		return "", err
	}

	val = u
	return u.Code, nil
}
