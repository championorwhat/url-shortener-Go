package service

import (
	"crypto/rand"
	"encoding/base64"
	"url-shortener/internal/model"
	"url-shortener/internal/repository"
)

type URLService struct {
	repo *repository.URLRepository
}

func NewURLService(repo *repository.URLRepository) *URLService {
	return &URLService{repo: repo}
}

func generateShortCode() string {
	b := make([]byte, 6)
	rand.Read(b)
	return base64.RawURLEncoding.EncodeToString(b)
}

func (s *URLService) CreateShortURL(originalURL string) (*model.URL, error) {
	url := &model.URL{
		ShortCode:   generateShortCode(),
		OriginalURL: originalURL,
	}

	err := s.repo.Save(url)
	if err != nil {
		return nil, err
	}

	return url, nil
}

func (s *URLService) GetOriginalURL(code string) (*model.URL, error) {
	return s.repo.FindByCode(code)
}
