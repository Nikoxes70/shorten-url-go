package inmem

import (
	"fmt"

	"shorten-url-go/shorturl/shorturl"
)

type shorturlsRepository struct {
	ShortURLs map[string]shorturl.ShortURL
}

func NewShortURLRepository() *shorturlsRepository {
	return &shorturlsRepository{
		ShortURLs: map[string]shorturl.ShortURL{},
	}
}

func (r *shorturlsRepository) GetShortURL(id string) (shorturl.ShortURL, error) {
	shortURL, isOk := r.ShortURLs[id]
	if !isOk {
		return shorturl.ShortURL{}, fmt.Errorf("no shorturl found for id: %v", id)
	}
	return shortURL, nil
}

func (r *shorturlsRepository) GetShortURLs() []shorturl.ShortURL {
	var shortURLs []shorturl.ShortURL
	for _, shortURL := range r.ShortURLs {
		shortURLs = append(shortURLs, shortURL)
	}
	return shortURLs
}

func (r *shorturlsRepository) CreateShortURL(shortURL shorturl.ShortURL) error {
	if _, ok := r.ShortURLs[shortURL.ID]; ok {
		return fmt.Errorf("id already exist")
	}
	r.ShortURLs[shortURL.ID] = shortURL
	return nil
}

func (r *shorturlsRepository) IncrementShortURLCount(id string) error {
	if shortURL, ok := r.ShortURLs[id]; !ok {
		shortURL.Count++
		return nil
	}
	return fmt.Errorf("shortURL not found")
}

func (r *shorturlsRepository) UpdateShortURLRedirections(id string, redirections []shorturl.Redirection) error {
	if shortURL, ok := r.ShortURLs[id]; ok {
		shortURL.Redirections = redirections
		r.ShortURLs[id] = shortURL
		return nil
	}
	return fmt.Errorf("shortURL not found")
}
