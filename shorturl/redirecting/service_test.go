package redirecting

import (
	"testing"
	"time"

	"shorten-url-go/shorturl/shorturl"
)

const (
	mockSuccessID = "123"
	mockFailureID = "456"

	morning = "morning"
	noon    = "noon"
	evening = "evening"
)

type mockRepo struct {
	Called map[string]int
}

func (m mockRepo) GetShortURL(id string) (shorturl.ShortURL, error) {
	m.Called["SingleShortURL"]++
	switch id {
	case mockSuccessID:
		return mockShortURL(), nil
	case mockFailureID:
		return mockShortURLEmpty(), nil
	}
	return shorturl.ShortURL{}, nil
}

func (m mockRepo) IncrementShortURLCount(id string) error {
	m.Called["IncrementShortURLCount"]++
	return nil
}

func mocks() *mockRepo {
	mr := &mockRepo{Called: make(map[string]int)}
	return mr
}

func mockShortURL() shorturl.ShortURL {
	mockRedirections := mockRedirections()
	return shorturl.ShortURL{
		ID:           mockSuccessID,
		Redirections: mockRedirections,
		Count:        0,
	}
}

func mockShortURLEmpty() shorturl.ShortURL {
	return shorturl.ShortURL{
		ID:    mockFailureID,
		Count: 0,
	}
}

func mockRedirections() []shorturl.Redirection {
	return []shorturl.Redirection{{
		From: 1,
		To:   9,
		URL:  morning,
	}, {
		From: 10,
		To:   16,
		URL:  noon,
	}, {
		From: 17,
		To:   23,
		URL:  evening,
	}}
}

func TestService_RedirectSuccess(t *testing.T) {
	mr := mocks()
	s := NewService(mr)
	url, err := s.Redirect(mockSuccessID)
	if err != nil {
		t.Errorf("TestService_RedirectSuccess Redirect failed with error: %v\n", err)
	}

	expectedURL := ""
	switch hour := time.Now().Hour(); {
	case hour >= 0 && hour <= 9:
		expectedURL = morning
	case hour >= 10 && hour <= 16:
		expectedURL = noon
	case hour >= 17 && hour <= 23:
		expectedURL = evening
	}

	if url != expectedURL {
		t.Errorf("TestService_RedirectSuccess Redirect failed - wrong url was provided by the server - expected: %v, got: %v", expectedURL, url)
	}

	if mr.Called["SingleShortURL"] != 1 {
		t.Errorf("TestService_RedirectSuccess SingleShortURL was called %v times instead of 1\n", mr.Called["SingleShortURL"])
	}
}

func TestService_RedirectFailure(t *testing.T) {
	mr := mocks()
	s := NewService(mr)
	_, err := s.Redirect(mockFailureID)
	if err == nil {
		t.Errorf("TestService_RedirectFailure Redirect success when had to gail\n")
	}

	if mr.Called["SingleShortURL"] != 1 {
		t.Errorf("TestService_RedirectSuccess SingleShortURL was called %v times instead of 1\n", mr.Called["SingleShortURL"])
	}
}
