package managing

import (
	"fmt"
	"shorten-url-go/shorturl/shorturl"

	smid "shorten-url-go/foundation/id"
)

type IdType string

const (
	alphanumeric IdType = "alphanumeric"
	uuid         IdType = "uuid"
)

func (t IdType) IsValid() bool {
	return t == alphanumeric || t == uuid
}

type repo interface {
	CreateShortURL(shortURL shorturl.ShortURL) error
}

type service struct {
	repo repo
}

func NewService(repo repo) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateShortURL(redirections []shorturl.Redirection, idType IdType) (id string, err error) {
	switch idType {
	case alphanumeric:
		id, err = smid.GenerateID(3)
	case uuid:
		id, err = smid.GenerateUUID()
	}

	if err != nil {
		return "", fmt.Errorf("CreateShortURL failed to generate id - %v", err)
	}

	shortURL := shorturl.ShortURL{
		ID:           id,
		Redirections: redirections,
		Count:        0,
	}

	if err = s.repo.CreateShortURL(shortURL); err != nil {
		return "", fmt.Errorf("CreateShortURL.repo failed - %v", err)
	}

	return id, nil
}
