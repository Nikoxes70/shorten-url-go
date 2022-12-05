package healthchecker

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"shorten-url-go/shorturl/shorturl"
)

type repo interface {
	GetShortURLs() []shorturl.ShortURL
	UpdateShortURLRedirections(id string, redirections []shorturl.Redirection) error
}

type cheker struct {
	tickerTime time.Duration
	repo
}

func NewHealthChecker(d time.Duration, repo repo) cheker {
	return cheker{
		d,
		repo,
	}
}

func (c *cheker) Start(ctx context.Context) {
	t := time.NewTicker(c.tickerTime)
	run := true
	go func(c context.Context) {
		for {
			select {
			case <-c.Done():
				run = false
				return
			}
		}
	}(ctx)
	for {
		select {
		case <-t.C:
			if run {
				c.do()
			}
		case <-ctx.Done():
			return
		}
	}
}

func (c *cheker) do() {
	shortURLs := c.repo.GetShortURLs()

	for _, shortURL := range shortURLs {
		c.checkRedirections(shortURL.ID, shortURL.Redirections)
	}
}

func (c *cheker) checkRedirections(id string, redirections []shorturl.Redirection) {
	shouldUpdate := false

	var validRedirections []shorturl.Redirection
	for _, redirection := range redirections {
		if err := c.check(redirection.URL); err != nil {
			shouldUpdate = true
		} else {
			validRedirections = append(redirections, redirection)
		}
	}
	if shouldUpdate {
		if err := c.repo.UpdateShortURLRedirections(id, validRedirections); err != nil {
			fmt.Printf("repo.UpdateShortURLRedirections failed - %v\n", err)
		}
	}
}

func (c *cheker) check(url string) error {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode > 400 {
		return fmt.Errorf("failed to GET url - %v, err: %v", url, err)
	}
	return nil
}
