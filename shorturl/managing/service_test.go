package managing

import (
	"testing"

	"shorten-url-go/shorturl/shorturl"
)

const (
	idLen   = 6
	uuidLen = 36
)

type mockRepo struct {
	Called map[string]int
}

func (m mockRepo) CreateShortURL(shortURL shorturl.ShortURL) error {
	m.Called["CreateShortURL"]++
	return nil
}

func mocks() *mockRepo {
	mr := &mockRepo{Called: make(map[string]int)}
	return mr
}

func mockRedirections() []shorturl.Redirection {
	return []shorturl.Redirection{{
		From: 1,
		To:   2,
		URL:  "url",
	}}
}

func TestService_CreateShortURLWithID(t *testing.T) {
	mr := mocks()
	s := NewService(mr)
	redirections := mockRedirections()
	id, err := s.CreateShortURL(redirections, alphanumeric)
	if err != nil {
		t.Errorf("TestService_CreateShortURL CreateShortURL failed with error: %v\n", err)
	}

	if len(id) != idLen {
		t.Errorf("TestService_CreateShortURL CreateShortURL failed - wrong id len was created exected: %v, got: %v\n", idLen, len(id))
	}

	if mr.Called["CreateShortURL"] != 1 {
		t.Errorf("TestService_CreateShortURL CreateShortURL was called %v times instead of 1\n", mr.Called["CreateShortURL"])
	}
}

func TestService_CreateShortURLWithUUID(t *testing.T) {
	mr := mocks()
	s := NewService(mr)
	redirections := mockRedirections()
	id, err := s.CreateShortURL(redirections, uuid)
	if err != nil {
		t.Errorf("TestService_CreateShortURL CreateShortURL failed with error: %v\n", err)
	}

	if len(id) != uuidLen {
		t.Errorf("TestService_CreateShortURL CreateShortURL failed - wrong id len was created exected: %v, got: %v\n", uuidLen, len(id))
	}

	if mr.Called["CreateShortURL"] != 1 {
		t.Errorf("TestService_CreateShortURL CreateShortURL was called %v times instead of 1\n", mr.Called["CreateShortURL"])
	}
}
