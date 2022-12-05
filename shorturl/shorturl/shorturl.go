package shorturl

import (
	"fmt"
	"time"
)

type ShortURL struct {
	ID           string
	Redirections []Redirection
	Count        int64
}

type Redirection struct {
	From int64
	To   int64
	URL  string
}

func (s *ShortURL) FetchUrl() (string, error) {
	currentHour := int64(time.Now().Hour())

	for _, redirection := range s.Redirections {
		if currentHour >= redirection.From && currentHour <= redirection.To {
			return redirection.URL, nil
		}
	}

	return "", fmt.Errorf("not found")
}
