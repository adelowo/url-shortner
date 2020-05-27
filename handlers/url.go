package handlers

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/adelowo/url-shortner/datastore"
	"github.com/adelowo/url-shortner/graph/model"
)

func CreateURL(ctx context.Context, db datastore.Store, s string) (*model.URL, error) {

	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}

	code, err := createCode()
	if err != nil {
		return nil, err
	}

	value := &datastore.URL{
		Code:        code,
		RedirectURL: u,
		CreatedAt:   time.Now(),
	}

	var id = int(value.ID)
	var url = value.RedirectURL.String()
	var created = value.CreatedAt.String()

	code, err = db.Create(ctx, value)
	if err != nil {
		return nil, err
	}

	return &model.URL{
		ID:          &id,
		Code:        &code,
		CreatedAt:   &created,
		RedirectURL: &url,
	}, nil
}

func createCode() (string, error) {
	d := make([]byte, 3)

	_, err := rand.Read(d)
	if err != nil {
		return "", errors.New("cannot create code")
	}

	return fmt.Sprintf("%x", string(d)), nil
}
