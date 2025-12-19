package repository

import (
	"database/sql"
	"errors"
	"url-shortener/internal/model"
)

type URLRepository struct {
	db *sql.DB
}

func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) Save(url *model.URL) error {
	query := `
		INSERT INTO urls (short_code, original_url)
		VALUES (?, ?)
		RETURNING id, created_at
	`

	return r.db.QueryRow(
		query,
		url.ShortCode,
		url.OriginalURL,
	).Scan(&url.ID, &url.CreatedAt)
}

func (r *URLRepository) FindByCode(code string) (*model.URL, error) {
	query := `SELECT id, short_code, original_url, created_at FROM urls WHERE short_code = ?`
	row := r.db.QueryRow(query, code)

	var url model.URL
	err := row.Scan(&url.ID, &url.ShortCode, &url.OriginalURL, &url.CreatedAt)
	if err != nil {
		return nil, errors.New("url not found")
	}

	return &url, nil
}
