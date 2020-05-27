package postgres

import "github.com/adelowo/url-shortner/datastore"

var _ datastore.Store = (*postgresDriver)(nil)
