package redirecting

import (
	"fmt"

	"shorten-url-go/shorturl/shorturl"
)

type repo interface {
	GetShortURL(id string) (shorturl.ShortURL, error)
	IncrementShortURLCount(id string) error
}

type service struct {
	repo repo
}

func NewService(repo repo) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Redirect(id string) (string, error) {
	shortURl, err := s.repo.GetShortURL(id)
	if err != nil {
		return "", fmt.Errorf("CreateShortURL.repo failed for id: %v - %v", id, err)
	}

	url, err := shortURl.FetchUrl()
	if err != nil {
		return "", fmt.Errorf("CreateShortURL.shortURl.fetchUrl failed for id: %v - %v", id, err)
	}

	err = s.repo.IncrementShortURLCount(id)
	if err != nil {
		fmt.Printf("failed to increment shorturl count for id: %v - %v\n", id, err)
	}

	return url, nil
}
